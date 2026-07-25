package module_02

import (
	"fmt"
	"io"
)

func Range(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Range over Slice")
	fmt.Fprintln(w, "----------------")

	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Fprintf(w, "index=%d value=%d\n", index, value)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Range Value is a Copy")
	fmt.Fprintln(w, "---------------------")

	for _, value := range numbers {
		value *= 10
	}

	fmt.Fprintln(w, "Slice after modification attempt:", numbers)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Modify the Original Slice")
	fmt.Fprintln(w, "-------------------------")

	for i := range numbers {
		numbers[i] *= 10
	}

	fmt.Fprintln(w, "Slice after index modification:", numbers)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Range over Map")
	fmt.Fprintln(w, "--------------")

	ages := map[string]int{
		"Alice": 20,
		"Bob":   25,
		"Carol": 30,
	}

	for name, age := range ages {
		fmt.Fprintf(w, "%s -> %d\n", name, age)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Range over String")
	fmt.Fprintln(w, "-----------------")

	text := "Go 世界"

	for index, r := range text {
		fmt.Fprintf(w, "byte=%d rune=%q unicode=%U\n", index, r, r)
	}

	return nil
}
