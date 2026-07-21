package main

import ( "fmt"
)


func makeRateLimiter(limit int) func() bool {
	count := 0
	return func() bool {
		if count < limit {
			count++
			return true
		}
		return false
	}
}

func main() {
	limiter := makeRateLimiter(3)
	for i := 0; i < 5; i++ {
		fmt.Printf("%v ",limiter())
	}
	fmt.Println()
}