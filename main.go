package main

import (
	"fmt"
	"os"

	"github.com/REZ0AN/go-verse/cli"
	_ "github.com/REZ0AN/go-verse/hello"
	_ "github.com/REZ0AN/go-verse/module_01"
	_ "github.com/REZ0AN/go-verse/module_02"
	// _ "github.com/REZ0AN/go-verse/module_03"
	// _ "github.com/REZ0AN/go-verse/module_04"
	// _ "github.com/REZ0AN/go-verse/module_05"
)

func main() {
	if err := cli.Run(
		os.Args,
		os.Stdin,
		os.Stdout,
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
