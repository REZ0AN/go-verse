package module_02

import (
	"fmt"
	"io"
)

type Resource struct {
	ID int
}

func (r Resource) Close(w io.Writer) {
	fmt.Fprintf(w, "Closing resource %d\n", r.ID)
}

func (r Resource) ImmediateClose(w io.Writer) {
	fmt.Fprintf(w, "Immediately Closing Resource %d\n", r.ID)
	defer r.Close(w)
}
func DeferInLoop(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Opening resources")
	fmt.Fprintln(w, "-----------------")

	for i := 1; i <= 3; i++ {
		r := Resource{ID: i}

		defer r.Close(w)

		fmt.Fprintf(w, "Using resource %d\n", r.ID)
	}

	fmt.Fprintf(w, "\n\n")

	fmt.Fprintln(w, "When you need immediate close after using the resource: ")
	fmt.Fprintln(w, "----------------------------------------------------------")
	for i := 1; i <= 3; i++ {
		r := Resource{ID: i}

		r.ImmediateClose(w)

		fmt.Fprintf(w, "Using resource %d\n", r.ID)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "All resources remain open until the function returns.")
	return nil
}
