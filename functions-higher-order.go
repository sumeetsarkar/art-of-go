package main

import (
	"fmt"
)

// Regular multiplier function
func mul(a, b int) int {
	return a * b
}

//----------------------------------------

// Higher Order Functions
func doubleit(fn func(a, b int) int) func(num int) int {
	wrapper := func(num int) int {
		// Due to property of closures, fn is accessible inside wrapper
		return fn(2, num)
	}
	return wrapper
}

// Higher Order Functions
func times(fn func(a, b int) int) func(multiplier int) func(num int) int {
	wrapper := func(multiplier int) func(num int) int {
		wrapper2 := func(num int) int {
			// Due to property of closures,
			// fn and multiplier is accessible inside wrapper2
			return fn(multiplier, num)
		}
		return wrapper2
	}
	return wrapper
}

//----------------------------------------

// Custom types declared, to demonstrate Functiona Programming in main()

type Employee struct {
	fname  string
	lname  string
	salary float64
}

type EmployeeList struct {
	list []Employee
}

func (el EmployeeList) filter(fn func(e Employee) bool) EmployeeList {
	filteredList := []Employee{}
	for _, v := range el.list {
		if fn(v) == true {
			filteredList = append(filteredList, v)
		}
	}
	return EmployeeList{
		filteredList,
	}
}

func (el EmployeeList) printall() {
	for _, v := range el.list {
		fmt.Println(v.fname, v.lname, v.salary)
	}
}

//----------------------------------------

func main() {
	// Anonymous Functions
	// Declare the function, just like a type declaration,
	// for functions the signature needs to be mentioned, param list with type and return type
	var greetUser func(name string) string

	// Define the function
	greetUser = func(name string) string {
		return ("Hello " + name)
	}

	// Invoke the function
	result := greetUser("Sumeet")
	fmt.Println(result)

	// or simply use shorthand
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(10, 20))

	twice := doubleit(mul)  // returns wrapper function, which can take num to double it
	fmt.Println(twice(100)) // 200
	fmt.Println(twice(200)) // 400

	times5 := times(mul)(5)
	fmt.Println(times5(100)) // 500

	times8 := times(mul)(8)
	fmt.Println(times8(100)) // 800

	// Functional Programming

	listOfEmployees := EmployeeList{
		list: []Employee{
			Employee{"Harry", "Potter", 250000},
			Employee{"John", "Wick", 300000},
			Employee{"Agent", "47", 500000},
		},
	}
	fmt.Println("\nPrint all employees........")
	listOfEmployees.printall()

	fmt.Println("\nPrint all employees........ with salary > 400000")
	resultEmployees := listOfEmployees.filter(func(e Employee) bool {
		return e.salary > 400000
	})
	resultEmployees.printall()
}
