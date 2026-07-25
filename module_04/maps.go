package module_04

import (
	"fmt"
	"io"
)

func Maps(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Creating a Map")
	fmt.Fprintln(w, "--------------")

	ages := map[string]int{
		"Alice": 20,
		"Bob":   25,
	}

	fmt.Fprintln(w, ages)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Creating with make")
	fmt.Fprintln(w, "------------------")

	scores := make(map[string]int)

	scores["Math"] = 90
	scores["English"] = 85

	fmt.Fprintln(w, scores)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Reading Values")
	fmt.Fprintln(w, "--------------")

	fmt.Fprintf(w, "Alice -> %d\n", ages["Alice"])

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Checking if a Key Exists")
	fmt.Fprintln(w, "------------------------")

	age, ok := ages["Charlie"]

	fmt.Fprintf(w, "Age = %d, Exists = %t\n", age, ok)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Deleting a Key")
	fmt.Fprintln(w, "--------------")

	delete(ages, "Bob")

	fmt.Fprintln(w, ages)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Iterating over a Map")
	fmt.Fprintln(w, "--------------------")

	for name, age := range ages {
		fmt.Fprintf(w, "%s -> %d\n", name, age)
	}

	return nil
}
