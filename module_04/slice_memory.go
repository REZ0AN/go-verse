package module_04

import (
	"fmt"
	"io"
)

func modify(w io.Writer, numbers []int) {
	fmt.Fprintf(
		w,
		"Inside modify (before): len=%d cap=%d %v\n",
		len(numbers),
		cap(numbers),
		numbers,
	)

	// Modifies the shared underlying array.
	numbers[0] = 999

	// May allocate a new underlying array.
	numbers = append(numbers, 100)

	fmt.Fprintf(
		w,
		"Inside modify (after) : len=%d cap=%d %v\n",
		len(numbers),
		cap(numbers),
		numbers,
	)
}

func SliceMemory(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Slices Share an Underlying Array")
	fmt.Fprintln(w, "--------------------------------")

	original := []int{1, 2, 3}

	fmt.Fprintf(
		w,
		"Before modify: len=%d cap=%d %v\n",
		len(original),
		cap(original),
		original,
	)

	modify(w, original)

	fmt.Fprintf(
		w,
		"After modify : len=%d cap=%d %v\n",
		len(original),
		cap(original),
		original,
	)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Only the element modification is visible.")
	fmt.Fprintln(w, "The appended element is not.")

	return nil
}
