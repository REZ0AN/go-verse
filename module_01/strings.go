package module_01

import (
	"fmt"
	"io"
)

// Strings demonstrates common string operations.
func Strings(_ io.Reader, w io.Writer) error {

	greeting := "Hello, World!"

	fmt.Fprintln(w, "Basic Operations")
	fmt.Fprintln(w, "----------------")

	fmt.Fprintf(w, "Original      : %q\n", greeting)
	fmt.Fprintf(w, "Length (bytes): %d\n", len(greeting))
	fmt.Fprintf(w, "First byte    : %q\n", greeting[0])
	fmt.Fprintf(w, "Slice         : %q\n", greeting[7:12])

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Concatenation")
	fmt.Fprintln(w, "-------------")

	name := "Alice"
	message := greeting + " My name is " + name + "."

	fmt.Fprintln(w, message)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Immutability")
	fmt.Fprintln(w, "------------")

	// Strings cannot be modified in-place.
	// greeting[0] = 'J' // Compile-time error

	newGreeting := "J" + greeting[1:]

	fmt.Fprintf(w, "Original : %q\n", greeting)
	fmt.Fprintf(w, "Modified : %q\n", newGreeting)

	return nil
}
