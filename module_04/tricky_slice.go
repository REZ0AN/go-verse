package main

import "fmt"

func modify(s []int) {
    s[0] = 999        // mutate existing element
	// s and original share the SAME underlying array (only the header was copied)
    // writing through s[0] mutates that shared array directly
    // underlying array is now: [999, 2, 3]
    // original ALREADY sees this change, since it points at the same array
    s = append(s, 100) // append — may or may not reallocate
	// len(s) == cap(s) == 3 → NO ROOM → append triggers REALLOCATION
    // Go allocates a brand NEW, larger array, copies [999, 2, 3] into it,
    // then appends 100 → new array: [999, 2, 3, 100]
    // s is REASSIGNED to point at this NEW array (s now has len=4, cap=6ish)
    // but this reassignment only changes s's LOCAL COPY of the header —
    // original's header, back in main(), was never touched by this line at all

	fmt.Println("Inside modify:", cap(s))

	/// **IMPORTANT**
	/// if a function needs to append in a way the caller should see, it must return the new slice, and the caller must reassign — original = modify(original) — exactly the same discipline as s = append(s, x)
}

func main() {
    original := []int{1, 2, 3}
	fmt.Println("Before modify:", cap(original))
    modify(original) // original's slice HEADER (ptr, len=3, cap=3) is COPIED into s
	fmt.Println("After modify:", cap(original))
    fmt.Println(original)
}