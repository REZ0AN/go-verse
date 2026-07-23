package module_05

import (
	"fmt"
	"strconv"
)

type Animal interface {
	Speak() string
	Run() string
}

type Dog struct {
	name string
}

type Lion struct {
	name string
}

func (d Dog) Speak() string {
	return d.name + " speaks woof"
}

func (d Dog) Run() string {
	return d.name + " is running"
}

func (l Lion) Speak() string {
	return l.name + " speaks Rooaaarr"
}

func (l Lion) Run() string {
	return l.name + " is running"
}

type Profile struct {
	Name string
	Age  int
}

func (p Profile) Describe() string {
	return "A " + strconv.Itoa(p.Age) + " aged person named " + p.Name
}

type Admin struct {
	Profile // embedded "Not extended"
	role    string
}

func AutoCheckTypeAssertions(a interface{}) {

	// a.(type) a special syntax only valid inside a switch statement
	switch v := a.(type) {
	case Dog:
		fmt.Println("It's a Dog named ", v.name)
	case Lion:
		fmt.Println("It's a Lion named ", v.name)
	default:
		fmt.Println("Unknown Animal")
	}
}
func Interface_Examples() {
	var a Animal
	a = Dog{
		"Bruno",
	}
	fmt.Println(a.Speak())
	fmt.Println(a.Run())

	fmt.Println("Empty Interfaces")
	// empty interfaces
	var x interface{}
	x = 5
	fmt.Println("x: ", x)
	x = "hello"
	fmt.Println("x: ", x)
	x = Dog{
		name: "Shiku",
	}
	// fmt.Println(x.Run()) // compiler knows by statically type is that interface{} doesn't has any methods
	// because every type implements the empty interface

	// but we have a solution for this
	// by type assertion we can explicitly tell compilere about the concrete type
	// dog := x.(Dog) // I assert x holds Dog - panics if wrong
	// better way to handle is using two-return form

	fmt.Println("Type Assertion")
	dog, ok := x.(Dog)
	if ok {
		fmt.Println(dog.Speak())
	} else {
		fmt.Println("Undefined Animal")
	}

	// better way to handle this
	AutoCheckTypeAssertions(x)
	x = Lion{
		name: "Mustafa",
	}
	AutoCheckTypeAssertions(x)
	x = 5
	AutoCheckTypeAssertions(x)

	ap := Admin{
		Profile{
			"Abir",
			25,
		},
		"Admin",
	} // example of composition GO promotes methods and fields to Admin from Profile implicitly
	// Admin never substituable for a Profile type

	fmt.Println(ap.Describe())
}
