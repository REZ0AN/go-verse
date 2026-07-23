package module_02

import "fmt"

func classify(n int) string {
	// Implement the FizzBuzz logic using a switch statement if/else replacement
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

func classifyIE(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}

}

func getDayPrefix(day int) string {

	switch day {
	case 1:
		return "Mon"
	case 2:
		return "Tue"
	case 3:
		return "Wed"
	case 4:
		return "Thu"
	case 5:
		return "Fri"
	case 6:
		return "Sat"
	case 7:
		return "Sun"
	default:
		return "Invalid day"
	}
}

const (
	Monday int = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func getRemainder(nominator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return nominator % denominator, nil
}

func ControlFlow_Examples() {
	// Test the classify function with numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(classify(i))
	}

	// Test the classifyIE function with numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(classifyIE(i))
	}

	// Test the getDayPrefix function with enums for days of the week
	for day := Monday; day <= Sunday; day++ {
		fmt.Printf("Day %d: %s\n", day, getDayPrefix(day))
	}

	if remainder, err := getRemainder(10, 3); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remainder of 10 divided by 3 is:", remainder)
	}

	// will show an error with undefined variable remainder
	// fmt.Println("Remainder %d", remainder)

	if remainder, err := getRemainder(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Remainder of 10 divided by 0 is:", remainder)
	}
}
