package module_04

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_04",
	Description: "Collections and receiver gotchas i.e maps, slices, and value vs pointer receivers",
	Examples: map[string]cli.Example{
		"tricky_map": {
			Name:        "tricky_map",
			Description: "Why m[key].Field = x won't compile on maps",
			Run:         TrickyMap_Examples,
		},

		"tricky_slice": {
			Name:        "tricky_slice",
			Description: "Slice header copy semantics and reallocation on append",
			Run:         TrickySlice_Examples,
		},

		"tricky_slice_2": {
			Name:        "tricky_slice_2",
			Description: "Shared backing arrays between slices within capacity",
			Run:         TrickySlice2_Examples,
		},

		"value_vs_pointer_reciever": {
			Name:        "value_vs_pointer_reciever",
			Description: "Value vs pointer method receivers",
			Run:         ValueVsPointerReceiver_Examples,
		},
	},
}
