package module_02

import (
	"fmt"
	"io"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func dayName(day int) string {
	switch day {
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	case Sunday:
		return "Sunday"
	default:
		return "Invalid"
	}
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func Switch(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Switch with Value")
	fmt.Fprintln(w, "-----------------")

	for day := Monday; day <= Sunday; day++ {
		fmt.Fprintf(w, "%d -> %s\n", day, dayName(day))
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Switch without Value")
	fmt.Fprintln(w, "--------------------")

	scores := []int{95, 83, 74, 68, 40}

	for _, score := range scores {
		fmt.Fprintf(w, "%d -> %s\n", score, grade(score))
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Multiple Values")
	fmt.Fprintln(w, "---------------")

	ch := 'a'

	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Fprintln(w, "Vowel")
	default:
		fmt.Fprintln(w, "Consonant")
	}

	return nil
}
