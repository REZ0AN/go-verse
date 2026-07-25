package module_04

import (
	"fmt"
	"io"
)

type Counter struct {
	Value int
}

func MapValues(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Map Values are Not Addressable")
	fmt.Fprintln(w, "------------------------------")

	counters := map[string]Counter{
		"visits": {Value: 1},
	}

	fmt.Fprintln(w, "Before:", counters)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "// counters[\"visits\"].Value = 2")
	fmt.Fprintln(w, "// Compiler error:")
	fmt.Fprintln(w, "// cannot assign to struct field counters[\"visits\"].Value in map")

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Update the Value")
	fmt.Fprintln(w, "----------------")

	counter := counters["visits"]
	counter.Value = 2
	counters["visits"] = counter

	fmt.Fprintln(w, "After:", counters)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Using Pointer Values")
	fmt.Fprintln(w, "--------------------")

	pointerMap := map[string]*Counter{
		"visits": {Value: 10},
	}

	pointerMap["visits"].Value = 20

	fmt.Fprintln(w, pointerMap["visits"])

	return nil
}
