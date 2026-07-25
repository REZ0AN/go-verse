package hello

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "hello",
	Description: "Learn the basics of Go output",
	Examples: map[string]cli.Example{
		"hello-world01": {
			Name:        "hello-world01",
			Description: "Print a simple message",
			Run:         HelloWorld01,
		},
		"hello-world02": {
			Name:        "hello-world02",
			Description: "Variables in output",
			Run:         HelloWorld02,
		},
	},
}
