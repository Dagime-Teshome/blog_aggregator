package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"

	"github.com/google/uuid"
)

func FeedsList(s *State, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.Db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error fetching feeds: %v", err)
	}
	for indx, feed := range feeds {
		userName, err := getUserName(ctx, s, feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user :%v", err)
		}
		printFeedInfo(userName, feed, indx)
	}
	return nil
}

func getUserName(ctx context.Context, s *State, id uuid.UUID) (string, error) {
	user, err := s.Db.GetUserById(ctx, id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func printFeedInfo(user string, feed database.Feed, indx int) {
	fmt.Printf("----- printing feed %v \n", indx)
	fmt.Printf("Feed Name: %s \n", feed.Name)
	fmt.Printf("Feed Url: %s \n", feed.Url)
	fmt.Printf("Feed User: %s \n", user)
}
