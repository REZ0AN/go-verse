package module_02

import (
	"fmt"
)

func labelsExample() {
	fmt.Println("Start of function")

	// Label for the outer loop
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Println(i, j)
				continue OuterLoop // Continue the outer loop when j == 2
			}
		}
	}
}

func exampleWithGoto() {
	fmt.Println("Start of function")

	// Label for the outer loop
	i := 0
outerLoop:
	if i < 3 {
		fmt.Println(i)
		i += 1
		goto outerLoop
	}
}

func GotoLabels_Examples() {
	fmt.Println("Demonstrating labels with continue: ")
	labelsExample()
	fmt.Println("\nDemonstrating labels with goto: ")
	exampleWithGoto()
}
