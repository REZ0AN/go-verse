package module_01

import (
	"fmt"
	"io"
)

// DataTypes demonstrates Go's built-in data types.
func DataTypes(_ io.Reader, w io.Writer) error {
	var (
		age        int        = 25
		population uint64     = 8_200_000_000
		height     float64    = 1.75
		isStudent  bool       = true
		grade      byte       = 'A'
		initial    rune       = '世'
		number     complex128 = 3 + 4i
	)

	fmt.Fprintln(w, "Numeric Types")
	fmt.Fprintln(w, "-------------")
	fmt.Fprintf(w, "int        : %d\n", age)
	fmt.Fprintf(w, "uint64     : %d\n", population)
	fmt.Fprintf(w, "float64    : %.2f\n", height)
	fmt.Fprintf(w, "complex128 : %v\n", number)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Other Types")
	fmt.Fprintln(w, "-----------")
	fmt.Fprintf(w, "bool : %t\n", isStudent)
	fmt.Fprintf(w, "byte : %c (%d)\n", grade, grade)
	fmt.Fprintf(w, "rune : %c (%U)\n", initial, initial)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Runtime Types")
	fmt.Fprintln(w, "-------------")
	fmt.Fprintf(w, "age       -> %T\n", age)
	fmt.Fprintf(w, "population-> %T\n", population)
	fmt.Fprintf(w, "height    -> %T\n", height)
	fmt.Fprintf(w, "isStudent -> %T\n", isStudent)
	fmt.Fprintf(w, "grade     -> %T\n", grade)
	fmt.Fprintf(w, "initial   -> %T\n", initial)
	fmt.Fprintf(w, "number    -> %T\n", number)

	return nil
}
