package module_01

import (
	"fmt"
)

func Formatting_Examples() {
	name := "Go"
	version := 1.23

	fmt.Println("Hello", name)                // space-separated, newline
	fmt.Printf("Go version: %.2f\n", version) // formatted, like C's printf
	s := fmt.Sprintf("v%.2f", version)        // formatted INTO a string, no print
	fmt.Println("Formatted string:", s)       // print the formatted string
	// common verbs
	// %d   integer
	// %f   float
	// %s   string
	// %v   "default" format for any value (great for debugging)
	// %+v  default format, but includes struct field names
	// %T   the TYPE of the value (extremely useful for debugging)
	// %t   boolean
}
