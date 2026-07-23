package module_05

import (
	"encoding/json"
	"fmt"
)

// only visible to this file
type cat struct {
	name  string
	beard string
}

func (c cat) Speak() string {
	return c.name + " is speaking meow"
}

// after providing compile time error
func (c cat) Run() string {
	return c.name + " is running"
}

func (c cat) Info() {
	fmt.Println("NAME: ", c.name)
	fmt.Println("Beard: ", c.beard)
}

// / named struct type
type album struct {
	title   string
	name    string
	authors []string
}

// / named struct type
type album1 struct {
	title string
	name  string
}

type album2 struct {
	title string
	name  string
}

type Response struct {
	Name   string   `json:"name"`
	Skills []string `json:"skills"`
}

func Struct_Examples() {
	var a1 = album{
		"ABC",
		"Programming like ABC",
		[]string{"JK", "PK"},
	}
	fmt.Println(a1)
	var a2 = album{
		"BCA",
		"Play Bayblade",
		[]string{"Ryuk", "SK"},
	}
	a1 = a2
	fmt.Println(a1)

	b1 := album1{
		"ABV",
		"wel",
	}
	fmt.Println(b1)
	b2 := album2{
		"KKK",
		"AKL",
	}
	fmt.Println(b2)

	// b1 = b2 // though internal structure is similar but names different
	// but we can typecast
	b1 = album1(b2) // why? as the internal structure is same thats why
	fmt.Println(b1)

	// anonymous struct
	a := struct {
		name string
		role string
	}{
		"Abir",
		"dev",
	}
	c := struct {
		name string
		role string
	}{
		"Lorin",
		"dev",
	}

	fmt.Println(a)
	b := &a

	fmt.Println(a, b)

	// why this works directly as they are anonymous and internal structure is same
	c = a

	fmt.Println(c)

	var json_n = Response{
		"Rezoan Ahmed Abir",
		[]string{"Python", "GO", "JavaScript", "C++"},
	}
	fmt.Println(json_n)
	json_, _ := json.Marshal(json_n)
	fmt.Println(string(json_))

	// Example using contract from another file
	var ck Animal
	ck = cat{
		name:  "Ginger",
		beard: "Ginger Cat",
	}

	// accessing the Animals contract functions
	fmt.Println(ck.Run())
	fmt.Println(ck.Speak())
	// but can't access the Info
	// fmt.Println(ck.Info()) // because it's not available in the Animal contract
	var cc cat = cat{
		name:  "Futu",
		beard: "Deshi",
	}

	// now you will see cc has all three why? concrete type
	fmt.Println(cc.Run())
	fmt.Println(cc.Speak())
	cc.Info()
}
