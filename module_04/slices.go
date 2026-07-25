package module_04

import (
	"fmt"
	"io"
)

func Slices(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Slice Declaration")
	fmt.Fprintln(w, "-----------------")

	var numbers []int

	fmt.Fprintln(w, "Zero value:", numbers)
	fmt.Fprintf(w, "len = %d, cap = %d\n", len(numbers), cap(numbers))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Slice Literal")
	fmt.Fprintln(w, "-------------")

	primes := []int{2, 3, 5, 7, 11}

	fmt.Fprintln(w, "Primes:", primes)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Creating with make")
	fmt.Fprintln(w, "------------------")

	values := make([]int, 3)

	fmt.Fprintln(w, "Values:", values)
	fmt.Fprintf(w, "len = %d, cap = %d\n", len(values), cap(values))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Creating with Length and Capacity")
	fmt.Fprintln(w, "---------------------------------")

	buffer := make([]int, 2, 5)

	fmt.Fprintln(w, "Buffer:", buffer)
	fmt.Fprintf(w, "len = %d, cap = %d\n", len(buffer), cap(buffer))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Slicing")
	fmt.Fprintln(w, "--------")

	sub := primes[1:4]

	fmt.Fprintln(w, "Original:", primes)
	fmt.Fprintln(w, "Slice   :", sub)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Iterating over a Slice")
	fmt.Fprintln(w, "----------------------")

	for index, value := range primes {
		fmt.Fprintf(w, "Index %d -> %d\n", index, value)
	}

	return nil
}
