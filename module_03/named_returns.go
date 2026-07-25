package module_03

import (
	"fmt"
	"io"
)

func divModNamed(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b

	return
}

func NamedReturns(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Named Return Values")
	fmt.Fprintln(w, "-------------------")

	quotient, remainder := divModNamed(17, 5)

	fmt.Fprintf(w, "17 / 5 = %d\n", quotient)
	fmt.Fprintf(w, "17 %% 5 = %d\n", remainder)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Named Returns are Zero Initialized")
	fmt.Fprintln(w, "----------------------------------")

	q, r := divModNamed(20, 6)

	fmt.Fprintf(w, "20 / 6 = %d\n", q)
	fmt.Fprintf(w, "20 %% 6 = %d\n", r)

	return nil
}
