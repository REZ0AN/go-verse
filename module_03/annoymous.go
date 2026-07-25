package module_03

import (
	"fmt"
	"io"
)

func Anonymous(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Anonymous Function")
	fmt.Fprintln(w, "------------------")

	result := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Fprintf(w, "10 + 20 = %d\n", result)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Assign Anonymous Function to a Variable")
	fmt.Fprintln(w, "---------------------------------------")

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Fprintf(w, "4 × 5 = %d\n", multiply(4, 5))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Functions are Values")
	fmt.Fprintln(w, "--------------------")

	var operation func(int, int) int

	operation = multiply

	fmt.Fprintf(w, "6 × 7 = %d\n", operation(6, 7))

	return nil
}
