package module_04

import (
	"fmt"
	"io"
)

func SliceAppend(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Appending Elements")
	fmt.Fprintln(w, "------------------")

	numbers := []int{1, 2, 3}

	fmt.Fprintf(w, "Before: %v (len=%d cap=%d)\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 4)

	fmt.Fprintf(w, "After : %v (len=%d cap=%d)\n", numbers, len(numbers), cap(numbers))

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Appending Multiple Elements")
	fmt.Fprintln(w, "---------------------------")

	numbers = append(numbers, 5, 6)

	fmt.Fprintf(w, "%v\n", numbers)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Appending Another Slice")
	fmt.Fprintln(w, "-----------------------")

	more := []int{7, 8, 9}

	numbers = append(numbers, more...)

	fmt.Fprintf(w, "%v\n", numbers)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Capacity Growth")
	fmt.Fprintln(w, "---------------")

	values := make([]int, 0, 2)

	for i := 1; i <= 5; i++ {
		values = append(values, i)
		fmt.Fprintf(
			w,
			"append(%d) -> len=%d cap=%d\n",
			i,
			len(values),
			cap(values),
		)
	}

	return nil
}
