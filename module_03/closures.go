package module_03

import (
	"fmt"
	"io"
)

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func fibonacciExp() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func Closures(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Counter Closure")
	fmt.Fprintln(w, "---------------")

	counter := makeCounter()

	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, "%d ", counter())
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Each Closure Has Its Own State")
	fmt.Fprintln(w, "------------------------------")

	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Fprintf(w, "counter1 -> %d\n", counter1())
	fmt.Fprintf(w, "counter1 -> %d\n", counter1())
	fmt.Fprintf(w, "counter2 -> %d\n", counter2())

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Fibonacci Closure")
	fmt.Fprintln(w, "-----------------")

	next := fibonacciExp()

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%d ", next())
	}

	fmt.Fprintln(w)

	return nil
}
