package module_03

import (
	"fmt"
	"io"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func Recursion(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Factorial")
	fmt.Fprintln(w, "---------")

	for i := 0; i <= 5; i++ {
		fmt.Fprintf(w, "%d! = %d\n", i, factorial(i))
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Recursive Fibonacci")
	fmt.Fprintln(w, "-------------------")

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%d ", fibonacci(i))
	}

	fmt.Fprintln(w)

	return nil
}
