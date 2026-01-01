package cli

import (
	"fmt"
	"go/gator/internal/config"
)

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

func (c *Commands) Register(name string, f func(*State, Command) error) {
	// add the command and the handler function to the commands map
	c.ComMap[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	// run the command that is sent using the args from the command argument
	value, ok := c.ComMap[cmd.Name]
	if !ok {
		return fmt.Errorf("Command doesn't exist")
	}
	if err := value(s, cmd); err != nil {
		return err
	}
	return nil
}
