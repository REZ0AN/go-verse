package module_02

import (
	"fmt"
	"io"
)

func recoverExample(w io.Writer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "Recovered from panic: %v\n", r)
		}
	}()

	fmt.Fprintln(w, "Before panic")

	panic("something went wrong")
}

func Recover(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Recover")
	fmt.Fprintln(w, "-------")

	recoverExample(w)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Program continues after recovery.")

	return nil
}
