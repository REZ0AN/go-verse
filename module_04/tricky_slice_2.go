package module_04

import "fmt"

func TrickySlice2_Examples() {
	s1 := make([]int, 2, 4) // len=2, cap=4
	s1[0], s1[1] = 1, 2

	s2 := append(s1, 3) // still within cap, no reallocation
	s2[0] = 999

	fmt.Println(s1)
	fmt.Println(s2)
}
