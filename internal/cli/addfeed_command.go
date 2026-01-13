package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func AddFeed(s *State, cmd Command) error {
	ctx := context.Background()
	if len(cmd.Args) < 1 {
		return fmt.Errorf("not enough arguments")
	}
	user, err := getUser(ctx, s)
	if err != nil {
		return err
	}
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
	fmt.Println(dbFeed)
	return nil
}

func getUser(ctx context.Context, s *State) (*database.User, error) {
	name := s.Config.Current_user_name
	user, err := s.Db.GetUser(ctx, name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
