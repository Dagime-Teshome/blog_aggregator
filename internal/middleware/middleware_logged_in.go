package middleware

import (
	"context"
	"fmt"

	"github.com/Dagime-Teshome/blog_aggregator/internal/cli"
	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
)

func LoggedInMiddleWare(handler func(s *cli.State, cmd cli.Command, user database.User) error) func(s *cli.State, cmd cli.Command) error {

	return func(s *cli.State, cmd cli.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error fetching user %w", err)
		}
		if err := handler(s, cmd, user); err != nil {
			return err
		}
		return nil
	}

}
