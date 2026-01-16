package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func Follow(s *State, cmd Command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("command needs url")
	}
	url := cmd.Args[0]
	userName := s.Config.Current_user_name
	ctx := context.Background()
	feed, err := s.Db.GetFeedByUrl(ctx, url)

	if err != nil {
		return fmt.Errorf("couldn't get feed %w", err)
	}

	user, err := s.Db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("error fetching user %w", err)
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
