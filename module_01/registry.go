package module_01

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_01",
	Description: "Go Fundamentals i.e variables, datatypes, operators, io, formatting, strings",
	Examples: map[string]cli.Example{
		"variables": {
			Name:        "variables",
			Description: "Variable declarations",
			Run:         Variable_Examples,
		},

		"datatypes": {
			Name:        "datatypes",
			Description: "Primitive data types",
			Run:         DataType_Examples,
		},

		"operators": {
			Name:        "operators",
			Description: "Arithmetic operators",
			Run:         Operator_Examples,
		},

		"formatting": {
			Name:        "formatting",
			Description: "fmt package examples",
			Run:         Formatting_Examples,
		},

		"io": {
			Name:        "io",
			Description: "Input Output examples",
			Run:         Io_Examples,
		},

		"strings": {
			Name:        "strings",
			Description: "String manipulation",
			Run:         String_Examples,
		},
	},
}
