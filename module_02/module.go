package module_02

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_02",
	Description: "Control Flow i.e if/else, switch, defer, panic, recover, goto and labels",
	Examples: map[string]cli.Example{
		"if": {
			Name:        "if",
			Description: "If, else and if initializer",
			Run:         IfStatements,
		},
		"switch": {
			Name:        "switch",
			Description: "Switch statements",
			Run:         Switch,
		},
		"loops": {
			Name:        "loops",
			Description: "For loops, break and continue",
			Run:         Loops,
		},
		"range": {
			Name:        "range",
			Description: "Iterate with range",
			Run:         Range,
		},
		"defer": {
			Name:        "defer",
			Description: "Deferred function calls",
			Run:         Defer,
		},
		"panic": {
			Name:        "panic",
			Description: "Create and understand panics",
			Run:         Panic,
		},
		"recover": {
			Name:        "recover",
			Description: "Recover from panics",
			Run:         Recover,
		},
		"goto": {
			Name:        "goto",
			Description: "Jump to labeled statements",
			Run:         Goto,
		},
		"fallthrough": {
			Name:        "fallthrough",
			Description: "Use the fallthrough statement",
			Run:         Fallthrough,
		},
		"defer_in_loop": {
			Name:        "defer_in_loop",
			Description: "Understand defer inside loops",
			Run:         DeferInLoop,
		},
	},
}
