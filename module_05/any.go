package module_05

import (
	"fmt"
	"io"
)

func Any(_ io.Reader, w io.Writer) error {
	var x interface{}

	x = Dog{}
	fmt.Fprintf(w, " %T \n\n", x)
	x = 5
	fmt.Fprintf(w, " %T \n\n", x)

	x = Women{}
	fmt.Fprintf(w, " %T \n\n", x)
	return nil
}
