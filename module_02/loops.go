package module_02

import (
	"fmt"
	"io"
)

func Loops(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Classic for Loop")
	fmt.Fprintln(w, "----------------")

	for i := 1; i <= 5; i++ {
		fmt.Fprintf(w, "i = %d\n", i)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "While-style Loop")
	fmt.Fprintln(w, "----------------")

	n := 1

	for n <= 5 {
		fmt.Fprintf(w, "n = %d\n", n)
		n++
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Infinite Loop")
	fmt.Fprintln(w, "-------------")

	count := 1

	for {
		fmt.Fprintf(w, "count = %d\n", count)

		if count == 5 {
			break
		}

		count++
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "break")
	fmt.Fprintln(w, "-----")

	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}

		fmt.Fprintf(w, "%d ", i)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "continue")
	fmt.Fprintln(w, "--------")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Fprintf(w, "%d ", i)
	}
	fmt.Fprintln(w)

	return nil
}
