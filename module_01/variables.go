package module_01

import (
	"fmt"
	"io"
)

// Variables demonstrates different ways to declare variables.
func Variables(_ io.Reader, w io.Writer) error {

	// Explicit type
	var language string = "Go"

	// Type inference
	var version = 1.25

	// Short variable declaration
	isOpenSource := true

	fmt.Fprintf(w, "Language: %s\n", language)
	fmt.Fprintf(w, "Version : %f\n", version)
	fmt.Fprintf(w, "Open     : %t\n", isOpenSource)

	fmt.Fprintln(w)

	// Multiple declarations
	var (
		author   = "Google"
		released = 2009
	)

	fmt.Fprintf(w, "Author   : %s\n", author)
	fmt.Fprintf(w, "Released : %d\n", released)

	return nil
}
