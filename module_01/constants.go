package module_01

import (
	"fmt"
	"io"
)

// Constants demonstrates constant values and iota.
func Constants(_ io.Reader, w io.Writer) error {

	// Untyped constants
	const pi = 3.14159
	const language = "Go"

	fmt.Fprintf(w, "Pi       : %.5f\n", pi)
	fmt.Fprintf(w, "Language : %s\n\n", language)

	// Typed constant
	const maxUsers int = 100

	fmt.Fprintf(w, "Max Users: %d\n\n", maxUsers)

	// iota
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Fprintln(w, "Weekdays:")
	fmt.Fprintf(w, "Sunday    = %d\n", Sunday)
	fmt.Fprintf(w, "Monday    = %d\n", Monday)
	fmt.Fprintf(w, "Tuesday   = %d\n", Tuesday)
	fmt.Fprintf(w, "Wednesday = %d\n", Wednesday)
	fmt.Fprintf(w, "Thursday  = %d\n", Thursday)
	fmt.Fprintf(w, "Friday    = %d\n", Friday)
	fmt.Fprintf(w, "Saturday  = %d\n", Saturday)

	return nil
}
