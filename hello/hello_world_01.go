package hello

import (
	"fmt"
	"io"
)

func HelloWorld01(_ io.Reader, w io.Writer) error {
	fmt.Fprintln(w, "Hello, Go")
	return nil
}
