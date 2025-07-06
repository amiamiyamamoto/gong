package cmd

import "fmt"

// HelpCommand shows help information
func HelpCommand(args []string) error {
	fmt.Println("Gong CLI Application")
	fmt.Println()
	fmt.Println("Usage: gong <command> [arguments]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  help      Show this help message")
	fmt.Println("  version   Show version information")
	fmt.Println()
	fmt.Println("Use 'gong <command> --help' for more information about a command.")
	return nil
}

// VersionCommand shows version information
func VersionCommand(args []string) error {
	fmt.Println("gong version 1.0.0")
	fmt.Println("Built with Go")
	return nil
}
