package module_01

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Input demonstrates reading input from an io.Reader.
func Input(r io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Reading a Word")
	fmt.Fprintln(w, "--------------")

	var name string

	fmt.Fprint(w, "Enter your name: ")

	if _, err := fmt.Fscan(r, &name); err != nil {
		return err
	}

	fmt.Fprintf(w, "Hello, %s!\n", name)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Reading a Full Line")
	fmt.Fprintln(w, "-------------------")

	// Consume the remaining newline left by Fscan.
	reader := bufio.NewReader(r)

	if _, err := reader.ReadString('\n'); err != nil && err != io.EOF {
		return err
	}

	fmt.Fprint(w, "Enter your favorite programming language: ")

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}

	line = strings.TrimSpace(line)

	fmt.Fprintf(w, "You entered: %q\n", line)

	return nil
}
