package cli

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}
	ctx := context.Background()
	fetchedUser, err := s.Db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user doesn't exist: %w", err)
	}
	if err := s.Config.SetUser(fetchedUser.Name); err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", cmd.Args[0])
	return nil
}
