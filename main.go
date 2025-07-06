package main

import (
	"fmt"
	"os"

	"github.com/amiamiyamamoto/gong/internal/cmd"
)

func main() {
	// Create command registry
	registry := cmd.NewRegistry()

	// Register commands
	registry.Register(&cmd.Command{
		Name:        "help",
		Description: "Show help information",
		Handler:     cmd.HelpCommand,
	})

	registry.Register(&cmd.Command{
		Name:        "version",
		Description: "Show version information",
		Handler:     cmd.VersionCommand,
	})

	// Parse command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: gong <command> [arguments]")
		fmt.Println("Run 'gong help' for more information")
		os.Exit(1)
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	// Execute command
	if err := registry.Execute(commandName, args); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Run 'gong help' for usage information")
		os.Exit(1)
	}
}
