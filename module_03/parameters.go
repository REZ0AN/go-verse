package module_03

import (
	"fmt"
	"io"
)

func introduce(name string, age int) {
	fmt.Fprintf(io.Discard, "") // avoid unused import if modified later
	fmt.Printf("")
	_ = name
	_ = age
}

func rectangleArea(width, height int) int {
	return width * height
}

func Parameters(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Single Parameters")
	fmt.Fprintln(w, "-----------------")

	name := "Alice"
	age := 25

	fmt.Fprintf(w, "Name: %s, Age: %d\n", name, age)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Grouped Parameters")
	fmt.Fprintln(w, "------------------")

	width := 10
	height := 5

	area := rectangleArea(width, height)

	fmt.Fprintf(w, "Rectangle: %d x %d\n", width, height)
	fmt.Fprintf(w, "Area: %d\n", area)

	return nil
}
