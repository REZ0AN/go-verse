package cli

import (
	"fmt"
	"io"
)

func PrintGlobalHelp(w io.Writer) {
	fmt.Fprintln(w, "Go Verse - A Go Learning Playground")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  go run . <module> <example>")
	fmt.Fprintln(w, "  go run . <module> --help")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "Available Modules:")
	fmt.Fprintln(w)

	for _, module := range AllModules() {
		fmt.Fprintf(w, "  %-15s %s\n", module.Name, module.Description)
	}
}

func PrintModuleHelp(w io.Writer, module Module) {
	fmt.Fprintf(w, "Module: %s\n", module.Name)
	fmt.Fprintf(w, "Description: %s\n\n", module.Description)

	fmt.Fprintln(w, "Usage:")
	fmt.Fprintf(w, "  go run . %s <example>\n", module.Name)
	fmt.Fprintf(w, "  go run . %s --help\n\n", module.Name)

	fmt.Fprintln(w, "Available Examples:")
	fmt.Fprintln(w)

	for _, example := range AllExamples(module) {
		fmt.Fprintf(w, "  %-20s %s\n", example.Name, example.Description)
	}
}
