package module_05

import (
	"fmt"
	"io"
)

type Walker interface {
	Walk(w io.Writer) string
}

type Wolf struct{}

func (*Wolf) Walk(_ io.Writer) string {
	return "Wolf is Running"
}

func ImplicitImplementation(_ io.Reader, w io.Writer) error {
	var wolP Wolf = Wolf{}
	var wal Walker = &wolP // this will work
	var wol Wolf

	// wal = wol this won't work
	fmt.Fprintln(w, wal.Walk(w)) // this will work
	fmt.Fprintln(w, wol.Walk(w)) // this works because compiler does some thing to convert it into (&wol).Walk()
	return nil
}
