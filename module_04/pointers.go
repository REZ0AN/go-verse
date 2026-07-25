package module_04

import (
	"fmt"
	"io"
)

func Pointers(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Pointer Declaration")
	fmt.Fprintln(w, "-------------------")

	x := 10

	var p *int = &x

	fmt.Fprintf(w, "x = %d\n", x)
	fmt.Fprintf(w, "p = %p\n", p)
	fmt.Fprintf(w, "*p = %d\n", *p)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Modify Through a Pointer")
	fmt.Fprintln(w, "------------------------")

	*p = 20

	fmt.Fprintf(w, "x = %d\n", x)
	fmt.Fprintf(w, "*p = %d\n", *p)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Pointer to a Struct")
	fmt.Fprintln(w, "-------------------")

	person := Person{
		Name: "Alice",
		Age:  25,
	}

	ptr := &person

	ptr.Age = 26
	ptr.Name = "Bob"

	fmt.Fprintln(w, person)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Pointers can be nil")
	fmt.Fprintln(w, "-------------------")

	var number *int

	fmt.Fprintf(w, "number == nil -> %t\n", number == nil)

	return nil
}
