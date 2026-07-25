package module_01

import (
	"fmt"
	"io"
)

// Formatting demonstrates Go's formatting functions and common verbs.
func Formatting(_ io.Reader, w io.Writer) error {

	name := "Go"
	version := 1.25

	fmt.Fprintln(w, "Print Functions")
	fmt.Fprintln(w, "---------------")

	fmt.Fprintln(w, "fmt.Fprintln:", "Hello", name)
	fmt.Fprintf(w, "fmt.Fprintf : Hello %s %.2f\n", name, version)

	message := fmt.Sprintf("%s %d", "Version", 125)
	fmt.Fprintf(w, "fmt.Sprintf: %q\n", message)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Common Formatting Verbs")
	fmt.Fprintln(w, "-----------------------")

	fmt.Fprintf(w, "%%d -> %d\n", 42)
	fmt.Fprintf(w, "%%f -> %.2f\n", 3.14159)
	fmt.Fprintf(w, "%%s -> %s\n", "Go")
	fmt.Fprintf(w, "%%t -> %t\n", true)
	fmt.Fprintf(w, "%%v -> %v\n", []int{1, 2, 3})
	fmt.Fprintf(w, "%%T -> %T\n", version)
	fmt.Fprintf(w, "%%q -> %q\n", "Hello")

	return nil
}
