package cli

import "fmt"

func PrintGlobalHelp() {
	fmt.Println("Go Verse - A Go Learning Playground")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("  go run . <module> <example>")
	fmt.Println("  go run . <module> --help")
	fmt.Println()

	fmt.Println("Available Modules:")
	fmt.Println()

	for _, module := range Modules {
		fmt.Printf("  %-15s %s\n", module.Name, module.Description)
	}
}

func PrintModuleHelp(module Module) {
	fmt.Printf("Helper for module -> %s\n", module.Name)
	fmt.Println()

	fmt.Println(module.Description)
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Printf("  go run . %s <example>\n", module.Name)
	fmt.Printf("  go run . %s --help\n", module.Name)
	fmt.Println()

	fmt.Println("Available Examples:")
	fmt.Println()

	for _, example := range module.Examples {
		fmt.Printf("  %-15s %s\n", example.Name, example.Description)
	}
}
