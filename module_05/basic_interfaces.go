package module_05

import (
	"fmt"
	"io"
)

type Speaker interface {
	Speak(w io.Writer)
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak(w io.Writer) {
	fmt.Fprintln(w, "Woof")
}

func (c Cat) Speak(w io.Writer) {
	fmt.Fprintln(w, "Meow")
}
func MakeSpeak(s Speaker, w io.Writer) {
	s.Speak(w)
}

func BasicInterfaces(r io.Reader, w io.Writer) error {

	MakeSpeak(Dog{}, w)
	MakeSpeak(Cat{}, w)
	return nil
}
