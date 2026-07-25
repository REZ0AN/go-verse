package module_03

import (
	"fmt"
	"io"
)

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func HigherOrder(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Passing Functions as Arguments")
	fmt.Fprintln(w, "------------------------------")

	fmt.Fprintf(w, "10 + 20 = %d\n", applyOperation(10, 20, add))
	fmt.Fprintf(w, "10 × 20 = %d\n", applyOperation(10, 20, multiply))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Using an Anonymous Function")
	fmt.Fprintln(w, "---------------------------")

	result := applyOperation(20, 5, func(a, b int) int {
		return a - b
	})

	fmt.Fprintf(w, "20 - 5 = %d\n", result)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Returning a Function")
	fmt.Fprintln(w, "--------------------")

	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Fprintf(w, "Double 8 = %d\n", double(8))
	fmt.Fprintf(w, "Triple 8 = %d\n", triple(8))

	return nil
}
