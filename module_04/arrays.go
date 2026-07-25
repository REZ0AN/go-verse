package module_04

import (
	"fmt"
	"io"
)

func Arrays(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Array Declaration")
	fmt.Fprintln(w, "-----------------")

	var numbers [5]int

	fmt.Fprintln(w, "Zero value:", numbers)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Array Literal")
	fmt.Fprintln(w, "-------------")

	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Fprintln(w, "Primes:", primes)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Length")
	fmt.Fprintln(w, "------")

	fmt.Fprintf(w, "len(primes) = %d\n", len(primes))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Arrays are Value Types")
	fmt.Fprintln(w, "----------------------")

	copyArray := primes
	copyArray[0] = 100

	fmt.Fprintln(w, "Original :", primes)
	fmt.Fprintln(w, "Copy     :", copyArray)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Iterating over an Array")
	fmt.Fprintln(w, "-----------------------")

	for index, value := range primes {
		fmt.Fprintf(w, "Index %d -> %d\n", index, value)
	}

	return nil
}
