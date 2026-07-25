package module_04

import (
	"fmt"
	"io"
)

type Person struct {
	Name string
	Age  int
}

func Structs(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Struct Declaration")
	fmt.Fprintln(w, "------------------")

	var person Person

	fmt.Fprintln(w, person)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Struct Literal")
	fmt.Fprintln(w, "--------------")

	alice := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Fprintln(w, alice)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Accessing Fields")
	fmt.Fprintln(w, "----------------")

	fmt.Fprintf(w, "Name: %s\n", alice.Name)
	fmt.Fprintf(w, "Age : %d\n", alice.Age)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Updating Fields")
	fmt.Fprintln(w, "---------------")

	alice.Age = 26

	fmt.Fprintln(w, alice)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Structs are Value Types")
	fmt.Fprintln(w, "-----------------------")

	copyPerson := alice
	copyPerson.Name = "Bob"

	fmt.Fprintln(w, "Original:", alice)
	fmt.Fprintln(w, "Copy    :", copyPerson)

	return nil
}
