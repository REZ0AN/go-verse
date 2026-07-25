package module_01

import (
	"fmt"
	"io"
)

// Hello demonstrates the simplest Go program.
func Hello(_ io.Reader, w io.Writer) error {
	fmt.Fprintln(w, "Hello, Go!")

	return nil
}
