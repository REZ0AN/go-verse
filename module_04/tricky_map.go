package module_04

import "fmt"

type Counter struct {
	Value int
}

func TrickyMap_Examples() {
	m := map[string]Counter{"a": {Value: 1}}
	// m["a"].Value = 99 // won't compile will show error as this m["a"] doesn't has real,stable address

	// sol
	c := m["a"]   // c has real address to be point on
	c.Value = 117 // updating value
	m["a"] = c    // a new address to point
	fmt.Println(m)
}
