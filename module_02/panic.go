package module_02

import (
	"fmt"
	"io"
)

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}

	return a / b
}

// Panic demonstrates how panic immediately stops normal execution.
func Panic(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Normal Execution")
	fmt.Fprintln(w, "----------------")

	fmt.Fprintf(w, "10 / 2 = %d\n", divide(10, 2))

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Panic")
	fmt.Fprintln(w, "-----")
	fmt.Fprintln(w, "Uncomment the line below to observe a panic.")

	divide(10, 0)

	fmt.Fprintln(w, "Program continues because the panic is commented out.")

	return nil
}
