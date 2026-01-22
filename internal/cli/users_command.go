package cli

import (
	"context"
	"fmt"
)

func GetUsers(s *State, cmd Command) error {
	ctx := context.Background()
	usersList, err := s.Db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Couldn't fetch user %w", err)
	}
	if len(usersList) == 0 {
		fmt.Println("no users in database")
		return nil
	}
	for _, value := range usersList {
		if value.Name == s.Config.Current_user_name {
			fmt.Printf("* %s (current)\n", value.Name)
			continue
		}
		fmt.Printf("* %s\n", value.Name)
	}
	return nil
}
