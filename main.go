package main

import (
	"github.com/REZ0AN/go-verse/cli"
	_ "github.com/REZ0AN/go-verse/hello"
	_ "github.com/REZ0AN/go-verse/module_01"
	_ "github.com/REZ0AN/go-verse/module_02"
	_ "github.com/REZ0AN/go-verse/module_03"
	_ "github.com/REZ0AN/go-verse/module_04"
)

func main() {
	cli.Run()
}
