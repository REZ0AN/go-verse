package module_01

import (
	"fmt"
	"io"
)

// ZeroValues demonstrates the default value assigned to variables.
func ZeroValues(_ io.Reader, w io.Writer) error {
	var (
		i  int
		f  float64
		b  bool
		s  string
		r  rune
		by byte
		p  *int
	)

	fmt.Fprintln(w, "Zero Values")
	fmt.Fprintln(w, "-----------")
	fmt.Fprintf(w, "int      : %d\n", i)
	fmt.Fprintf(w, "float64  : %f\n", f)
	fmt.Fprintf(w, "bool     : %t\n", b)
	fmt.Fprintf(w, "string   : %q\n", s)
	fmt.Fprintf(w, "rune     : %q (%d)\n", r, r)
	fmt.Fprintf(w, "byte     : %d\n", by)
	fmt.Fprintf(w, "pointer  : %v\n", p)

	return nil
}
