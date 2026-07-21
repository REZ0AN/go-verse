package main

import (

	"fmt"
)

func deferOrderExample() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Function body")
}

func deferArgTakingExample() {
	x := 5
	defer fmt.Println("Deferred:", x)
	x = 10
	fmt.Println("Function body:", x)
}

// Recover from panic
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b // panics if b == 0 (division by zero)
    return
}

// Panic example
func panicDivide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a/b
}

func cleanupExample(x int) {
	defer fmt.Println("Cleanup after processing", x)
}

func perIterationCleanupExample(vals []int) {
	for _, v := range vals {
	
		fmt.Println("Processing", v)
		cleanupExample(v)
	}
}
func process(vals []int) {
    defer fmt.Println("process finished")

    for _, v := range vals {
        defer fmt.Println("cleanup for", v)
        fmt.Println("processing", v)
    }
}



func main() {
	fmt.Println("Demonstrating defer order: ")
	deferOrderExample();
	fmt.Println("\nDemonstrating defer with argument taking: ")
	deferArgTakingExample();
	fmt.Println("\nDemonstrating panic: ")
	// Uncomment the following line to see panic in action
	// fmt.Println("Result of panicDivide(10, 0):", panicDivide(10, 0)) (program stops right here)
	fmt.Println("\nDemonstrating recover from panic: ")
	if result, err := safeDivide(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
    // Demonstrating cleanup with defer in a loop: ")
	fmt.Println("\nDemonstrating cleanup with defer in a loop: ")
	process([]int{1, 2, 3})

	// Demonstrating per-iteration cleanup with defer in a loop: ")
	fmt.Println("\nDemonstrating per-iteration cleanup with defer in a loop: ")
	perIterationCleanupExample([]int{4, 5, 6})
}