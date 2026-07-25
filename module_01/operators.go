package module_01

import (
	"fmt"
	"io"
)

// Operators demonstrates Go's operators.
func Operators(_ io.Reader, w io.Writer) error {

	a := 10
	b := 3

	fmt.Fprintln(w, "Arithmetic Operators")
	fmt.Fprintln(w, "--------------------")

	fmt.Fprintf(w, "%d + %d = %d\n", a, b, a+b)
	fmt.Fprintf(w, "%d - %d = %d\n", a, b, a-b)
	fmt.Fprintf(w, "%d * %d = %d\n", a, b, a*b)
	fmt.Fprintf(w, "%d / %d = %d\n", a, b, a/b)
	fmt.Fprintf(w, "%d %% %d = %d\n", a, b, a%b)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Comparison Operators")
	fmt.Fprintln(w, "--------------------")

	fmt.Fprintf(w, "%d == %d : %t\n", a, b, a == b)
	fmt.Fprintf(w, "%d != %d : %t\n", a, b, a != b)
	fmt.Fprintf(w, "%d > %d  : %t\n", a, b, a > b)
	fmt.Fprintf(w, "%d < %d  : %t\n", a, b, a < b)
	fmt.Fprintf(w, "%d >= %d : %t\n", a, b, a >= b)
	fmt.Fprintf(w, "%d <= %d : %t\n", a, b, a <= b)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Logical Operators")
	fmt.Fprintln(w, "-----------------")

	x := true
	y := false

	fmt.Fprintf(w, "x && y = %t\n", x && y)
	fmt.Fprintf(w, "x || y = %t\n", x || y)
	fmt.Fprintf(w, "!x     = %t\n", !x)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Assignment Operators")
	fmt.Fprintln(w, "--------------------")

	value := 10
	fmt.Fprintf(w, "Initial : %d\n", value)

	value += 5
	fmt.Fprintf(w, "+= 5    : %d\n", value)

	value -= 2
	fmt.Fprintf(w, "-= 2    : %d\n", value)

	value *= 3
	fmt.Fprintf(w, "*= 3    : %d\n", value)

	value /= 4
	fmt.Fprintf(w, "/= 4    : %d\n", value)

	value %= 2
	fmt.Fprintf(w, "%%= 2    : %d\n", value)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Bitwise Operators")
	fmt.Fprintln(w, "-----------------")

	c := 5 // 0101
	d := 3 // 0011

	fmt.Fprintf(w, "%04b &  %04b = %04b (%d)\n", c, d, c&d, c&d)
	fmt.Fprintf(w, "%04b |  %04b = %04b (%d)\n", c, d, c|d, c|d)
	fmt.Fprintf(w, "%04b ^  %04b = %04b (%d)\n", c, d, c^d, c^d)
	fmt.Fprintf(w, "%04b &^ %04b = %04b (%d)\n", c, d, c&^d, c&^d)
	fmt.Fprintf(w, "%04b << 1 = %04b (%d)\n", c, c<<1, c<<1)
	fmt.Fprintf(w, "%04b >> 1 = %04b (%d)\n", c, c>>1, c>>1)

	return nil
}
