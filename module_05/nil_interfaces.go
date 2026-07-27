package module_05

import (
	"fmt"
	"io"
)

func NilInterfaces(_ io.Reader, w io.Writer) error {
	var s Speaker
	fmt.Fprintln(w, s == nil)
	var d *Dog = nil // value nil, type Dog
	s = d
	fmt.Fprintln(w, s == nil)
	fmt.Fprintf(w, " %T \n", s)
	var dd Dog
	s = dd
	fmt.Fprintln(w, s == nil)
	fmt.Fprintf(w, " %T \n", s)

	return nil
}
