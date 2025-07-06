package cmd

import "fmt"

// HelpCommand shows help information
func HelpCommand(args []string) error {
	fmt.Println("🔔 Gong CLI Application")
	fmt.Println()
	fmt.Println("Usage: gong [command] [arguments]")
	fmt.Println()
	fmt.Println("When run without arguments, plays the gong sound.")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  (no args)  Play gong sound (default)")
	fmt.Println("  play       Play gong sound")
	fmt.Println("  help       Show this help message")
	fmt.Println("  version    Show version information")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  gong           # Play gong sound")
	fmt.Println("  gong play      # Play gong sound")
	fmt.Println("  gong help      # Show this help")
	return nil
}

// VersionCommand shows version information
func VersionCommand(args []string) error {
	fmt.Println("gong version 1.0.0")
	fmt.Println("Built with Go")
	return nil
}
