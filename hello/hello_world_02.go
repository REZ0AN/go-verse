package hello

import (
	"fmt"
	"io"
)

func HelloWorld02(_ io.Reader, w io.Writer) error {
	name := "Go"
	fmt.Fprintln(w, "Hello,", name)
	return nil
}
