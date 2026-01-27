package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Dagime-Teshome/blog_aggregator/internal/config"
	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
)

// State holds the application state including database and config
type State struct {
	Db     *database.Queries
	Config *config.Config
}

// Command represents a CLI command with its name and arguments
type Command struct {
	Name string
	Args []string
}

// CommandInfo stores metadata about a command for help display
type CommandInfo struct {
	Handler     func(*State, Command) error
	Description string
	Usage       string
}

// Commands is a registry for CLI commands
type Commands struct {
	handlers map[string]CommandInfo
}

// NewCommands creates a new command registry
func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]CommandInfo),
	}
}

// Register adds a command to the registry with metadata
func (c *Commands) Register(name string, handler func(*State, Command) error, description, usage string) {
	c.handlers[name] = CommandInfo{
		Handler:     handler,
		Description: description,
		Usage:       usage,
	}
}

// Run executes a command by name
func (c *Commands) Run(s *State, cmd Command) error {
	info, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s. Run 'help' for available commands", cmd.Name)
	}
	return info.Handler(s, cmd)
}

// Help returns help text for all registered commands
func (c *Commands) Help() string {
	var sb strings.Builder
	sb.WriteString("Gator - RSS Feed Aggregator\n\n")
	sb.WriteString("Available commands:\n")

	// Sort command names for consistent output
	names := make([]string, 0, len(c.handlers))
	for name := range c.handlers {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		info := c.handlers[name]
		sb.WriteString(fmt.Sprintf("  %-12s %s\n", name, info.Description))
		if info.Usage != "" {
			sb.WriteString(fmt.Sprintf("               Usage: %s\n", info.Usage))
		}
	}
	return sb.String()
}

// HandlerHelp displays help information
func HandlerHelp(commands *Commands) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		fmt.Print(commands.Help())
		return nil
	}
}
