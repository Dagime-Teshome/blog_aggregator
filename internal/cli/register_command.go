package cli

import (
	"context"
	"fmt"
	"go/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

// Ensure that a name was passed in the args.
// Create a new user in the database. It should have access to the CreateUser query through the state -> db struct.
// Pass context.Background() to the query to create an empty Context argument.
// Use the uuid.New() function to generate a new UUID for the user.
// created_at and updated_at should be the current time.
// Use the provided name.
// Exit with code 1 if a user with that name already exists.
// Set the current user in the config to the given name.
// Print a message that the user was created, and log the user's data to the console for your own debugging.
func Register(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Arguments fewer than expected")
	}
	user := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ctx := context.Background()
	created_user, err := s.Db.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	if err := s.Config.SetUser(created_user.Name); err != nil {
		return err
	}
	fmt.Printf("User %s is successfully created with data : %v\n", created_user.Name, created_user)
	return nil
}
