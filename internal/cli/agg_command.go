package cli

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
	feed "github.com/Dagime-Teshome/blog_aggregator/internal/rss_feed"

	"github.com/google/uuid"
)

func Agg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("not enough arguments")
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing request interval: %w", err)
	}

	ticker := time.NewTicker(timeBetweenReqs)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) {
	ctx := context.Background()
	fetchedFeed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Println("error getting next feed:", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.Db, fetchedFeed)
}

func scrapeFeed(db *database.Queries, db_feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), db_feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", db_feed.Name, err)
		return
	}

	feedData, err := feed.FetchFeed(context.Background(), db_feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", db_feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		createPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			FeedID:      db_feed.ID,
			PublishedAt: publishedAt,
			Title:       item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url: item.Link,
		}
		post, err := db.CreatePost(context.Background(), createPost)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("error creating post %v", post)
			continue
		}
		fmt.Println(post)
	}
	log.Printf("Feed %s collected, %v posts found", db_feed.Name, len(feedData.Channel.Item))
}
