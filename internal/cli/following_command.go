package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"

	"github.com/google/uuid"
)

func Following(s *State, cmd Command, user database.User) error {
	ctx := context.Background()
	feedsList, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds user follow: %w", err)
	}
	if len(feedsList) == 0 {
		fmt.Println("user doesn't follow any feeds")
		return nil
	}
	fmt.Println("---------------Feeds List---------------")
	for i, feed := range feedsList {
		feedName, err := getFeedName(ctx, s, feed.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("%v : %v \n", i, feedName)
	}
	return nil
}

func getFeedName(ctx context.Context, s *State, feeId uuid.UUID) (string, error) {
	feed, err := s.Db.GetFeedById(ctx, feeId)
	if err != nil {
		return "", err
	}
	return feed.Name, nil
}
