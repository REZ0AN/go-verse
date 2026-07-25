package module_04

import "github.com/REZ0AN/go-verse/cli"

var Module = cli.Module{
	Name:        "module_04",
	Description: "Collections and receiver gotchas i.e maps, slices, and value vs pointer receivers",
	Examples: map[string]cli.Example{
		"arrays": {
			Name:        "arrays",
			Description: "Arrays",
			Run:         Arrays,
		},

		"slices": {
			Name:        "slices",
			Description: "Slice basics",
			Run:         Slices,
		},
		"slice_append": {
			Name:        "slice_append",
			Description: "Appending to slices",
			Run:         SliceAppend,
		},
		"maps": {
			Name:        "maps",
			Description: "Maps",
			Run:         Maps,
		},
		"structs": {
			Name:        "structs",
			Description: "Structs",
			Run:         Structs,
		},
		"pointers": {
			Name:        "pointers",
			Description: "Pointers",
			Run:         Pointers,
		},
		"methods": {
			Name:        "methods",
			Description: "Methods",
			Run:         Methods,
		},
		"pointer_receivers": {
			Name:        "pointer_receivers",
			Description: "Value and pointer receivers",
			Run:         PointerReceivers,
		},
		"slice_memory": {
			Name:        "slice_memory",
			Description: "How slices share memory",
			Run:         SliceMemory,
		},
		"map_values": {
			Name:        "map_values",
			Description: "Map values are not addressable",
			Run:         MapValues,
		},
	},
}
