package module_02

import (
	"fmt"
	"io"
)

func classifySwitch(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func classifyIfElse(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}

func remainder(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a % b, nil
}

func IfStatements(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Simple if")
	fmt.Fprintln(w, "---------")

	if 10 > 5 {
		fmt.Fprintln(w, "10 is greater than 5")
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "if-else")
	fmt.Fprintln(w, "-------")

	fmt.Fprintln(w, classifyIfElse(3))
	fmt.Fprintln(w, classifyIfElse(5))
	fmt.Fprintln(w, classifyIfElse(15))
	fmt.Fprintln(w, classifyIfElse(7))

	fmt.Fprintln(w)
	fmt.Fprintln(w, "if initializer")
	fmt.Fprintln(w, "--------------")

	if rem, err := remainder(10, 3); err != nil {
		fmt.Fprintln(w, "Error:", err)
	} else {
		fmt.Fprintf(w, "10 %% 3 = %d\n", rem)
	}

	if rem, err := remainder(10, 0); err != nil {
		fmt.Fprintln(w, "Error:", err)
	} else {
		fmt.Fprintf(w, "10 %% 0 = %d\n", rem)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Switch as an if replacement")
	fmt.Fprintln(w, "--------------------------")

	classifySwitch(10)
	classifySwitch(97)
	classifySwitch(15)
	classifySwitch(6)

	return nil
}
