package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func Following(s *State, cmd Command) error {
	ctx := context.Background()
	userName := s.Config.Current_user_name
	user, err := s.Db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("Couldn't find user %w", err)
	}
	feedsList, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds user follow: %w", err)
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
