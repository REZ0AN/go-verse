package cli

import (
	"fmt"
	"io"
)

func Run(
	args []string,
	stdin io.Reader,
	stdout io.Writer,
) error {

	// go run .
	if len(args) == 1 {
		PrintGlobalHelp(stdout)
		return nil
	}

	// go run . --help
	if args[1] == "--help" || args[1] == "-h" {
		PrintGlobalHelp(stdout)
		return nil
	}

	moduleName := args[1]

	module, ok := LookupModule(moduleName)
	if !ok {
		return fmt.Errorf("unknown module %q", moduleName)
	}

	// go run . module_01
	if len(args) == 2 {
		PrintModuleHelp(stdout, module)
		return nil
	}

	// go run . module_01 --help
	if args[2] == "--help" || args[2] == "-h" {
		PrintModuleHelp(stdout, module)
		return nil
	}

	exampleName := args[2]

	example, ok := LookupExample(module, exampleName)
	if !ok {
		return fmt.Errorf("unknown example %q in module %q", exampleName, moduleName)
	}

	fmt.Fprintf(stdout, "Running %s -> %s\n\n", moduleName, exampleName)

	if err := example.Run(stdin, stdout); err != nil {
		return fmt.Errorf("%s/%s: %w", moduleName, exampleName, err)
	}

	return nil
}
