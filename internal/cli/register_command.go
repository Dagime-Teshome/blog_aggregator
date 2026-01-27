package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/Dagime-Teshome/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func Register(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("arguments fewer than expected")
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ctx := context.Background()
	createdUser, err := s.Db.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	if err := s.Config.SetUser(createdUser.Name); err != nil {
		return err
	}
	fmt.Printf("User %s is successfully created with data: %v\n", createdUser.Name, createdUser)
	return nil
}
