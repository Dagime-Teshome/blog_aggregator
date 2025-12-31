package cli

import "go/gator/internal/config"

type State struct {
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	ComMap map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {

}

func (c *Commands) run(s *State, cmd Command) error {
	return nil
}
