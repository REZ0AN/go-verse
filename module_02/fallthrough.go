package module_02

import (
	"fmt"
	"io"
)

func Fallthrough(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Switch without fallthrough")
	fmt.Fprintln(w, "-------------------------")

	n := 2

	switch n {
	case 1:
		fmt.Fprintln(w, "One")
	case 2:
		fmt.Fprintln(w, "Two")
	case 3:
		fmt.Fprintln(w, "Three")
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Switch with fallthrough")
	fmt.Fprintln(w, "-----------------------")

	switch n {
	case 1:
		fmt.Fprintln(w, "One")
	case 2:
		fmt.Fprintln(w, "Two")
		fallthrough
	case 3:
		fmt.Fprintln(w, "Three")
	default:
		fmt.Fprintln(w, "Done")
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w, "fallthrough ignores conditions")
	fmt.Fprintln(w, "------------------------------")

	x := 2

	switch {
	case x == 2:
		fmt.Fprintln(w, "x == 2")
		fallthrough
	case x == 100:
		fmt.Fprintln(w, "This still executes")
	default:
		fmt.Fprintln(w, "Default")
	}

	return nil
}
