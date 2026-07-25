package module_02

import (
	"fmt"
	"io"
)

func deferOrder(w io.Writer) {
	defer fmt.Fprintln(w, "First")
	defer fmt.Fprintln(w, "Second")
	defer fmt.Fprintln(w, "Third")

	fmt.Fprintln(w, "Function body")
}

func deferArguments(w io.Writer) {
	x := 5

	defer fmt.Fprintln(w, "Deferred value:", x)

	x = 10

	fmt.Fprintln(w, "Current value:", x)
}

func Defer(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Defer Execution Order")
	fmt.Fprintln(w, "---------------------")

	deferOrder(w)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Deferred Arguments")
	fmt.Fprintln(w, "------------------")

	deferArguments(w)

	return nil
}
