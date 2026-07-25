package module_03

import (
	"fmt"
	"io"
)

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func multiplyAndSum(multiplier int, numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number * multiplier
	}

	return total
}

func Variadic(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Variadic Function")
	fmt.Fprintln(w, "-----------------")

	total := sum(1, 2, 3, 4, 5)
	fmt.Fprintf(w, "Sum = %d\n", total)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Passing a Slice")
	fmt.Fprintln(w, "----------------")

	numbers := []int{10, 20, 30}

	total = sum(numbers...)

	fmt.Fprintf(w, "Sum = %d\n", total)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Fixed + Variadic Parameters")
	fmt.Fprintln(w, "---------------------------")

	result := multiplyAndSum(2, 1, 2, 3)

	fmt.Fprintf(w, "Result = %d\n", result)

	return nil
}
