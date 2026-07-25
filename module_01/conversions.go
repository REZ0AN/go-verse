package module_01

import (
	"fmt"
	"io"
	"strconv"
)

// Conversions demonstrates explicit type conversions in Go.
func Conversions(_ io.Reader, w io.Writer) error {

	// Numeric conversions
	var age int = 25
	var height float64 = float64(age)

	fmt.Fprintln(w, "Numeric Conversions")
	fmt.Fprintln(w, "-------------------")
	fmt.Fprintf(w, "int -> float64 : %d -> %.1f\n", age, height)

	var pi float64 = 3.14
	var whole int = int(pi)

	fmt.Fprintf(w, "float64 -> int : %.2f -> %d\n", pi, whole)

	// Integer to string
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Integer to String")
	fmt.Fprintln(w, "-----------------")

	number := 123

	incorrect := string(number) // ignore warnings!!
	fmt.Fprintf(w, "string(%d)      = %q (Unicode code point)\n", number, incorrect)

	correct := strconv.Itoa(number)
	fmt.Fprintf(w, "strconv.Itoa() = %q\n", correct)

	sprintf := fmt.Sprintf("%d", number)
	fmt.Fprintf(w, "fmt.Sprintf()  = %q\n", sprintf)

	// String to integer
	fmt.Fprintln(w)
	fmt.Fprintln(w, "String to Integer")
	fmt.Fprintln(w, "-----------------")

	text := "456"

	value, err := strconv.Atoi(text)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "strconv.Atoi(%q) = %d\n", text, value)

	return nil
}
