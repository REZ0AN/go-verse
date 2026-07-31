package module_05

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type Reader interface {
	Read(io.Reader, io.Writer)
}

type Writer interface {
	Write(io.Writer)
}

type ReadWriter interface {
	Reader
	Writer
}

type GetAndShow struct {
	name string
	age  uint
}

func (g *GetAndShow) Read(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)

	fmt.Fprintln(w, "Enter your name:")
	name, _ := reader.ReadString('\n')
	g.name = strings.TrimSpace(name)

	fmt.Fprintln(w, "Enter your age:")
	fmt.Fscan(reader, &g.age)
}

func (g *GetAndShow) Write(w io.Writer) {
	fmt.Fprintf(w, "\n\nGreetings %s\n\nYou were born in %d\n\n", g.name, (time.Now().Year() - int(g.age)))
}

func InterfaceComposition(r io.Reader, w io.Writer) error {
	var rw ReadWriter = &GetAndShow{}
	rw.Read(r, w)
	rw.Write(w)
	return nil
}
