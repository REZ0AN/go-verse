package module_05

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_05",
	Description: "Interfaces",
	Examples: map[string]cli.Example{
		"interface": {
			Name:        "interface",
			Description: "What, How, Why Interfaces",
			Run:         Interface_Examples,
		},

		"struct": {
			Name:        "struct",
			Description: "A detailed work through struct intuitions in GO",
			Run:         Struct_Examples,
		},
	},
}
