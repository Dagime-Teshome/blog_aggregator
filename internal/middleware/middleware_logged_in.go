package middleware

import (
	"context"
	"fmt"
	"go/gator/internal/cli"
	"go/gator/internal/database"
)

func LoggedInMiddleWare(handler func(s *cli.State, cmd cli.Command, user database.User) error) func(s *cli.State, cmd cli.Command) error {

	return func(s *cli.State, cmd cli.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return fmt.Errorf("error fetching user %w", err)
		}
		if err := handler(s, cmd, user); err != nil {
			return err
		}
		return nil
	}

}
