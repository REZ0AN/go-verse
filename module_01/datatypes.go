package module_01

import (
	"fmt"
)

func DataType_Examples() {

	// Numeric Data Types

	// int, uint (range depends on the platform architecture)
	// uint8, uint16, uint32, uint64 (range depends on the platform architecture)
	// int8, int16, int32, int64 (ranges int8 -> -128 to 127, int16 -> -32768 to 32767, int32 -> -2147483648 to 2147483647, int64 -> -9223372036854775808 to 9223372036854775807)
	// float32, float64 (ranges float32 -> ±1.18e-38 to ±3.4e38, float64 -> ±2.23e-308 to ±1.8e308)
	// Complex64, Complex128
	// Declare a variable of type int
	var age int = 25
	fmt.Println("Age:", age)

	// Declare a variable of type float64
	var height float64 = 5.9
	fmt.Println("Height:", height)

	// Declare a variable of type string
	var name string = "John Doe"
	fmt.Println("Name:", name)

	// Declare a variable of type bool
	var isStudent bool = true
	fmt.Println("Is Student:", isStudent)

	// Declare a variable of type rune
	var initial rune = 'J'
	fmt.Println("Initial:", initial)

	// Declare a variable of type byte
	var grade byte = 90
	fmt.Println("Grade:", grade)

	// type conversion
	var x int = 10
	var y float64 = float64(x) // convert int to float64
	fmt.Println("Converted value:", y)
	fmt.Printf("Type of x: %T\n", x) // print the type of x
	fmt.Printf("Type of y: %T\n", y) // print the type of y

}
