package module_03

import (
	"fmt"
	"io"
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

func RateLimiter(_ io.Reader, w io.Writer) error {

	fmt.Fprintln(w, "Simple Rate Limiter")
	fmt.Fprintln(w, "-------------------")

	limiter := makeRateLimiter(3)

	for i := 1; i <= 5; i++ {
		allowed := limiter()
		fmt.Fprintf(w, "Request %d -> %t\n", i, allowed)
	}

	fmt.Fprintln(w)

	fmt.Fprintln(w, "Independent Limiters")
	fmt.Fprintln(w, "--------------------")

	userA := makeRateLimiter(2)
	userB := makeRateLimiter(2)

	fmt.Fprintf(w, "User A: %t\n", userA())
	fmt.Fprintf(w, "User A: %t\n", userA())
	fmt.Fprintf(w, "User A: %t\n", userA())

	fmt.Fprintln(w)

	fmt.Fprintf(w, "User B: %t\n", userB())
	fmt.Fprintf(w, "User B: %t\n", userB())
	fmt.Fprintf(w, "User B: %t\n", userB())

	return nil
}
