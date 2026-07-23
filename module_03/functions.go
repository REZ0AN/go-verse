package module_03

import (
	"fmt"
)

// mutliple return values
func divmod(a, b int) (int, int) {
	d := a / b
	r := a % b
	return d, r

}

// named return values
func divmodNamed(a, b int) (d int, r int) {
	d = a / b
	r = a % b
	return
}

// variadic function
func sumXx(x int, a ...int) int {
	total := 0
	for _, v := range a {
		total += v * x
	}
	return total
}

// closure function
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// HOC function
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func Function_Examples() {
	fmt.Println("Demonstrating multiple return values:")
	d, r := divmod(10, 3)
	fmt.Printf("Div: %d, Mod: %d\n", d, r)

	fmt.Println("Demonstrating named return values:")
	d2, r2 := divmodNamed(30, 3)
	fmt.Printf("Div: %d, Mod: %d\n", d2, r2)

	fmt.Println("Demonstrating variadic function:")
	result := sumXx(4, 2, 4, 8)
	fmt.Printf("Result of sumXx: %d\n", result)

	// anonymous function
	fmt.Println("Demonstrating anonymous function:")
	fmt.Println("Result of anonymous function:", func(x, y int) int { return x * y }(3, 4))

	// closure function
	f := fibonacci()
	fmt.Println("Demonstrating closure function:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Println()

	// HOC function
	fmt.Println("Demonstrating higher-order function:")
	result = applyOperation(5, 3, func(x, y int) int { return x * y })
	fmt.Printf("Result of higher-order function: %d\n", result)
}
