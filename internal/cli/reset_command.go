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
	fmt.Print("User table reset \n")
	return nil
}
