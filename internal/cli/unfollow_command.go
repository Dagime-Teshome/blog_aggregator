package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"
)

func UnfollowFeed(s *State, cmd Command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("not enough command arguments")
	}
	ctx := context.Background()
	feed, err := getFeed(s, ctx, cmd.Args[0])
	if err != nil {
		return err
	}
	unfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.Db.DeleteFeedFollow(ctx, unfollowParams)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}
	fmt.Println("feed unfollowed successfully")
	return nil
}

func getFeed(s *State, ctx context.Context, url string) (*database.Feed, error) {
	feed, err := s.Db.GetFeedByUrl(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed %w", err)
	}
	return &feed, nil
}
