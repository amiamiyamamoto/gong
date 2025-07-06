package cmd

import "fmt"

// Command represents a CLI command
type Command struct {
	Name        string
	Description string
	Handler     func(args []string) error
}

// Registry holds all available commands
type Registry struct {
	commands map[string]*Command
}

// NewRegistry creates a new command registry
func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]*Command),
	}
}

// Register adds a new command to the registry
func (r *Registry) Register(cmd *Command) {
	r.commands[cmd.Name] = cmd
}

// Execute runs the specified command with given arguments
func (r *Registry) Execute(name string, args []string) error {
	cmd, exists := r.commands[name]
	if !exists {
		return fmt.Errorf("unknown command: %s", name)
	}
	return cmd.Handler(args)
}

// List returns all registered commands
func (r *Registry) List() map[string]*Command {
	return r.commands
}
