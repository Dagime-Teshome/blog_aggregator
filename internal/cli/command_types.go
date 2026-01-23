package cli

import (
	"fmt"

	"github.com/Dagime-Teshome/blog_aggregator/internal/config"
	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
)

type State struct {
	Db     *database.Queries
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
	c.ComMap[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	value, ok := c.ComMap[cmd.Name]
	if !ok {
		return fmt.Errorf("Command doesn't exist")
	}
	if err := value(s, cmd); err != nil {
		return err
	}
	return nil
}
