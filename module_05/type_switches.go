package module_05

import (
	"fmt"
	"io"
)

func (m Men) Speak(w io.Writer) {
	fmt.Fprintln(w, "Men")
}
func Describe(s Speaker, w io.Writer) {
	switch s.(type) {
	case Dog:
		fmt.Fprintln(w, "Dog")
	case Cat:
		fmt.Fprintln(w, "Cat")

	default:
		fmt.Fprintln(w, "No Matches Found")
	}
}
func TypeSwitches(_ io.Reader, w io.Writer) error {

	Describe(Dog{}, w)
	Describe(Cat{}, w)

	// must implement the interface if want to use type assertions and type switches
	Describe(Men{}, w)

	return nil
}
