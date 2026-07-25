package module_01

import (
	"fmt"
	"io"
	"unicode/utf8"
)

// Unicode demonstrates how Go represents strings using UTF-8.
func Unicode(_ io.Reader, w io.Writer) error {

	text := "Hello, 世界"

	fmt.Fprintln(w, "String")
	fmt.Fprintln(w, "------")

	fmt.Fprintf(w, "Value            : %q\n", text)
	fmt.Fprintf(w, "Length (bytes)   : %d\n", len(text))
	fmt.Fprintf(w, "Length (runes)   : %d\n", utf8.RuneCountInString(text))

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Iterating with range")
	fmt.Fprintln(w, "--------------------")

	for index, r := range text {
		fmt.Fprintf(
			w,
			"byte=%2d rune=%c unicode=%U\n",
			index,
			r,
			r,
		)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Converting to []rune")
	fmt.Fprintln(w, "--------------------")

	runes := []rune(text)

	fmt.Fprintf(w, "Rune slice length: %d\n", len(runes))

	for index, r := range runes {
		fmt.Fprintf(
			w,
			"index=%d rune=%c unicode=%U\n",
			index,
			r,
			r,
		)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "byte vs rune")
	fmt.Fprintln(w, "------------")

	var b byte = 'A'
	var r rune = '世'

	fmt.Fprintf(w, "byte : %c (%d)\n", b, b)
	fmt.Fprintf(w, "rune : %c (%U)\n", r, r)

	return nil
}
