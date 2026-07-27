package module_05

import (
	"fmt"
	"io"
)

func TypeAssertions(_ io.Reader, w io.Writer) error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(w, "Recovered from panic: %v\n\n", r)
		}
	}()

	var s Speaker = Dog{}

	// from interface values
	s.Speak(w)

	var d = s.(Dog)
	// from the actual concrete implementation
	d.Speak(w)

	// program panics
	// var c = s.(Cat)
	// fmt.Fprintf(w, " %T \n\n", c) // below code won't run

	// another form we can handle without panics
	cc, ok := s.(Cat)
	if ok {
		fmt.Fprintln(w, "It's a Dog")
	} else {
		fmt.Fprintln(w, "It's a Cat", cc)
	}

	return nil
}
