package cli

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("User name Required")
	}
	ctx := context.Background()
	feted_user, err := s.Db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("User doesn't exist %v", err)
	}
	err = s.Config.SetUser(feted_user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", cmd.Args[0])
	return nil
}
