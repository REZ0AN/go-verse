package module_05

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_05",
	Description: "Interfaces, implicit satisfaction, method sets, dynamic dispatch",
	Examples: map[string]cli.Example{
		"basic_interfaces": {
			Name:        "basic_interfaces",
			Description: "Declaring interfaces; satisfaction is implicit — no 'implements' keyword",
			Run:         BasicInterfaces,
		},
		"implicit_implementation": {
			Name:        "implicit_implementation",
			Description: "Method sets: pointer receivers exclude T; auto &x needs an addressable operand",
			Run:         ImplicitImplementation,
		},
		"interface_values": {
			Name:        "interface_values",
			Description: "The 2-word (itab, data) value; compile-time assertions; dispatch via itab.fun",
			Run:         InterfaceValues,
		},
		"nil_interfaces": {
			Name:        "nil_interfaces",
			Description: "Typed-nil trap: an interface is nil only when BOTH words are nil",
			Run:         NilInterfaces,
		},
		"type_assertions": {
			Name:        "type_assertions",
			Description: "Recovering the concrete type: x.(T) panics, x.(T) with comma-ok does not",
			Run:         TypeAssertions,
		},
		"type_switches": {
			Name:        "type_switches",
			Description: "switch x.(type): branching on dynamic type, plus the nil and default cases",
			Run:         TypeSwitches,
		},
		"any": {
			Name:        "any",
			Description: "any is an ALIAS for interface{} — identical type, and boxing still allocates",
			Run:         Any,
		},
		"interface_composition": {
			Name:        "interface_composition",
			Description: "New interface by embedding one or more existing ones, supports composition over inheritance, makes code easier to test, mock, satisfy, reusable",
			Run:         InterfaceComposition,
		},
	},
}
