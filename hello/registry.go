package hello

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "hello",
	Description: "Hello world from GO learning",
	Examples: map[string]cli.Example{
		"hello-world01": {
			Name:        "hello-world01",
			Description: "Just simple Print",
			Run:         Hello_World_Example,
		},
		"hello-world02": {
			Name:        "hello-world02",
			Description: "Tweak with args and a short hand declaration",
			Run:         Hello_World02_Example,
		},
	},
}
