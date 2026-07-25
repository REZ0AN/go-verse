package module_03

import (
	"fmt"
	"io"
)

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func addition(a, b int) int {
	return a + b
}

func Functions(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Function with No Return Value")
	fmt.Fprintln(w, "-----------------------------")

	greet("Go")

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Function with Return Value")
	fmt.Fprintln(w, "--------------------------")

	result := addition(10, 20)
	fmt.Fprintf(w, "10 + 20 = %d\n", result)

	return nil
}
