package cli

import (
	"fmt"
)

func Follow(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("command needs url")
	}
	// fix this command
	// feedParams := database.FollowFeedParams{
	// 	ID: uuid.New(),
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	FeedID: "getsomeid",
	// 	UserID: "getsomeuserid",
	// }
	return nil
}
