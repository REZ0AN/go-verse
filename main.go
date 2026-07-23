package main

import (
	"fmt"

	"github.com/REZ0AN/go-verse/module_01"
)

func main() {
	fmt.Println("Daemon Running!!")

	// MODULE-01 RUN one at a time
	// Variable Examples
	fmt.Println("Exploring Variable Examples: ")
	module_01.Variable_Examples()

	// Data Type Examples
	fmt.Println("Exploring Data Type Examples: ")
	module_01.DataType_Examples()

	// Operator Examples
	fmt.Println("Exploring Operator Examples: ")
	module_01.Operator_Examples()

	// Formatting Examples
	fmt.Println("Exploring Formatting Examples: ")
	module_01.Formatting_Examples()

	// Input Output Examples
	fmt.Println("Exploring IO Examples: ")
	module_01.Io_Examples()

	// String Examples
	fmt.Println("Exploring String Examples: ")
	module_01.String_Examples()

}
