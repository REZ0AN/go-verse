package module_02

import (
	"fmt"
	"io"
)

func Goto(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Basic goto")
	fmt.Fprintln(w, "----------")

	count := 1

Start:
	fmt.Fprintf(w, "count = %d\n", count)

	count++

	if count <= 5 {
		goto Start
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Skipping Code")
	fmt.Fprintln(w, "-------------")

	fmt.Fprintln(w, "Before goto")

	goto End

	fmt.Fprintln(w, "This line is skipped") // ignore warning!

End:
	fmt.Fprintln(w, "After label")

	return nil
}
