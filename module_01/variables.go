package main

import (
	"fmt"
)


func main() {
	// declaring  variables - Three ways
	var name string = "Go" // explicit type, explicit initialization
	var age = 10 // implicit type, explicit initialization
    isAdult := false // implicit type, implicit initialization
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Adult:", isAdult)

	// multiple variable declaration at one time
	var (
    nameAgain string = "Go"
    year int    = 2009
    stable bool = true
)

	fmt.Println("Name:", nameAgain)
	fmt.Println("Year:", year)
	fmt.Println("Stable:", stable)

	// constants
	const pi = 3.14
	const (
    StatusOK       = 200
    StatusNotFound = 404
	)

	var areaCircle float32 = pi * 10 * 10;
	fmt.Println("Area of Circle:", areaCircle)
	fmt.Printf("Type of Pi: %T\n", pi)
	fmt.Printf("Type of Area of Circle: %T\n", areaCircle)
	fmt.Printf("Status OK: %d, Status Not Found: %d\n", StatusOK, StatusNotFound)

	// enums in GO using iota
	type Weekday int
	const (
	Sunday  Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	)
	fmt.Println("Days of the week:")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

	// type of Sunday
	fmt.Printf("Type of Sunday: %T\n", Sunday)
}
