package module_03

import (
	"fmt"
	"io"
)

func divMod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func MultipleReturns(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Multiple Return Values")
	fmt.Fprintln(w, "----------------------")

	quotient, remainder := divMod(17, 5)

	fmt.Fprintf(w, "17 / 5 = %d\n", quotient)
	fmt.Fprintf(w, "17 %% 5 = %d\n", remainder)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Ignoring Return Values")
	fmt.Fprintln(w, "----------------------")

	quotient, _ = divMod(20, 3)

	fmt.Fprintf(w, "20 / 3 = %d\n", quotient)

	return nil
}
