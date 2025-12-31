package cli

import "fmt"

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects one argument but has received none")
	}
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", cmd.Args[0])
	return nil
}
