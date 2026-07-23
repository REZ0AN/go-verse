package module_02

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_02",
	Description: "Control Flow i.e if/else, switch, defer, panic, recover, goto and labels",
	Examples: map[string]cli.Example{
		"control_flow": {
			Name:        "control_flow",
			Description: "if/else, switch, and day-of-week classification",
			Run:         ControlFlow_Examples,
		},

		"exceptions_cleanup": {
			Name:        "exceptions_cleanup",
			Description: "Defer order, panic, recover, and cleanup patterns",
			Run:         ExceptionsCleanup_Examples,
		},

		"goto_labels": {
			Name:        "goto_labels",
			Description: "Labeled loops with continue and goto",
			Run:         GotoLabels_Examples,
		},
	},
}
