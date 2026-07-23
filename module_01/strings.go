package module_01

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func String_Examples() {
	// declare a string variable
	var greeting string = "Hello, World!"
	fmt.Println(greeting)

	// string concatenation
	var name string = "Alice"
	var message string = greeting + " My name is " + name + "."
	fmt.Println(message)

	// string length
	fmt.Println("Length of greeting:", len(greeting)) // gives the number of bytes in the string

	// string indexing
	fmt.Println("First character of greeting:", string(greeting[0])) // indexing starts at 0

	// string slicing
	fmt.Println("Slice of greeting:", greeting[7:12]) // prints "World"

	// strings are immutable, so we cannot change a character in a string directly
	// but we can create a new string based on the old one
	// greeting[0] = 'J' // this will cause a compile-time error

	newGreeting := "J" + greeting[1:] // replace 'H' with 'J'
	fmt.Println("New Greeting:", newGreeting)

	// byte vs rune
	var b byte = 'A'
	fmt.Println("Byte value:", b) // prints 65, the ASCII value of 'A'

	var r rune = '世'
	fmt.Println("Rune value:", r) // prints 19990, the Unicode code point for '世'

	var str string = "Hello, 世界"
	fmt.Println("String:", str)
	fmt.Println("Length of string in bytes:", len(str))                    // length in bytes
	fmt.Println("Length of string in runes:", utf8.RuneCountInString(str)) // length in runes (count of Unicode characters

	// will see index jump over the bytes of the multi-byte characters
	for i, r := range str {
		fmt.Printf("Character %d: %c (Unicode: %U)\n", i, r, r)
	}

	// converting string to []rune
	runes := []rune(str)
	fmt.Println("Runes:", runes)
	fmt.Println("Length of runes:", len(runes)) // length in runes
	for i, r := range runes {
		fmt.Printf("Rune %d: %c (Unicode: %U)\n", i, r, r)
	}

	// convert int to string
	var num int = 123
	strNum := string(num) // this will not work as expected, it will convert to the character with Unicode code point 123
	fmt.Println("String from int (incorrect):", strNum)

	// correct way to convert int to string
	strNumCorrect := fmt.Sprintf("%d", num)
	fmt.Println("String from int (correct):", strNumCorrect)

	// another
	strNumCorrect2 := strconv.Itoa(num)
	fmt.Println("String from int (correct using strconv):", strNumCorrect2)

	// convert string to int
	strToInt := "456"
	intValue, err := strconv.Atoi(strToInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("Converted int value:", intValue)
	}

}
