package module_04

import (
	"fmt"
	"io"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func Methods(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Calling Methods")
	fmt.Fprintln(w, "---------------")

	rect := Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Fprintf(w, "Area      = %.2f\n", rect.Area())
	fmt.Fprintf(w, "Perimeter = %.2f\n", rect.Perimeter())

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Methods on Different Values")
	fmt.Fprintln(w, "---------------------------")

	rect2 := Rectangle{
		Width:  8,
		Height: 3,
	}

	fmt.Fprintf(w, "Area = %.2f\n", rect2.Area())

	return nil
}
