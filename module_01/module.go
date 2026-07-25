package module_01

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_01",
	Description: "Go language fundamentals",
	Examples: map[string]cli.Example{

		"hello": {
			Name:        "hello",
			Description: "Print output",
			Run:         Hello,
		},

		"variables": {
			Name:        "variables",
			Description: "Declare and initialize variables",
			Run:         Variables,
		},

		"constants": {
			Name:        "constants",
			Description: "Constants and iota",
			Run:         Constants,
		},

		"zero-values": {
			Name:        "zero-values",
			Description: "Default values of Go types",
			Run:         ZeroValues,
		},

		"data-types": {
			Name:        "data-types",
			Description: "Built-in data types",
			Run:         DataTypes,
		},

		"conversions": {
			Name:        "conversions",
			Description: "Explicit type conversions",
			Run:         Conversions,
		},

		"formatting": {
			Name:        "formatting",
			Description: "Formatted output",
			Run:         Formatting,
		},

		"strings": {
			Name:        "strings",
			Description: "String operations",
			Run:         Strings,
		},

		"unicode": {
			Name:        "unicode",
			Description: "UTF-8, bytes, and runes",
			Run:         Unicode,
		},

		"operators": {
			Name:        "operators",
			Description: "Arithmetic, comparison, logical, and bitwise operators",
			Run:         Operators,
		},

		"input": {
			Name:        "input",
			Description: "Read input from an io.Reader",
			Run:         Input,
		},
	},
}
