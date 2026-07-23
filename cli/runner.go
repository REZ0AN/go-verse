package cli

import (
	"fmt"
	"os"
)

func Run() {
	args := os.Args

	// go run .
	if len(args) == 1 {
		PrintGlobalHelp()
		return
	}

	moduleName := args[1]

	module, ok := Modules[moduleName]
	if !ok {
		fmt.Printf("Unknown module: %s\n\n", moduleName)
		PrintGlobalHelp()
		return
	}

	// go run . module_01
	if len(args) == 2 {
		PrintModuleHelp(module)
		return
	}

	// go run . module_01 --help
	if args[2] == "--help" || args[2] == "-h" {
		PrintModuleHelp(module)
		return
	}

	exampleName := args[2]

	example, ok := module.Examples[exampleName]
	if !ok {
		fmt.Printf("Unknown example: %s\n\n", exampleName)
		PrintModuleHelp(module)
		return
	}

	fmt.Printf("Running %s -> %s Examples\n\n", module.Name, example.Name)

	example.Run()
}
