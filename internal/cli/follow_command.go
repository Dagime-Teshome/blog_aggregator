package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func Follow(s *State, cmd Command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("command needs url")
	}
	url := cmd.Args[0]
	ctx := context.Background()
	feed, err := s.Db.GetFeedByUrl(ctx, url)

	if err != nil {
		return fmt.Errorf("couldn't get feed %w", err)
	}

	feedParams := database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	followFeed, err := s.Db.FollowFeed(ctx, feedParams)
	fmt.Println(followFeed)
	return nil
}
