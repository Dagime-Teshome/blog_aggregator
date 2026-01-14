package cli

import (
	"context"
	"fmt"
)

func Reset(s *State, cmd Command) error {
	ctx := context.Background()

	if err := s.Db.ResetUserTable(ctx); err != nil {
		return err
	}
	if err := s.Db.ResetFeeds(ctx); err != nil {
		return err
	}
	fmt.Print("Tables reset \n")
	return nil
}
