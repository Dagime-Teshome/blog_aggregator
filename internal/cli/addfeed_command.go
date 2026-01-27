package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Dagime-Teshome/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func AddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("not enough arguments")
	}
	ctx := context.Background()
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dbFeed, err := s.Db.CreateFeed(ctx, feed)
	if err != nil {
		return err
	}
	err = autoFeedFollow(s, ctx, user.ID, dbFeed.ID)
	if err != nil {
		return err
	}
	fmt.Printf("feed %s created with url %s added to feed list\n", dbFeed.Name, dbFeed.Url)
	return nil
}

func autoFeedFollow(s *State, ctx context.Context, userId uuid.UUID, feedId uuid.UUID) error {

	feedParams := database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feedId,
		UserID:    userId,
	}
	followFeed, err := s.Db.FollowFeed(ctx, feedParams)
	if err != nil {
		return err
	}
	fmt.Printf(" %s user following %s feed\n", s.Config.CurrentUserName, followFeed.FeedName)
	return nil
}
