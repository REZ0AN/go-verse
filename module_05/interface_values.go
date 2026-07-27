package module_05

import (
	"fmt"
	"io"
)

type Viewer interface {
	See()
}

type Men struct{}
type Women struct{}

func (Men) See()   {}
func (Women) See() {}

func InterfaceValues(_ io.Reader, _ io.Writer) error {
	var v Viewer

	v = Men{}
	fmt.Printf("%T\n", v)

	v = Women{}
	fmt.Printf("%T\n", v)
	return nil
}
