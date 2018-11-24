package main

import (
	"fmt"
)

// Methods are functions, with:
// special receiver type declared between func keyword and method name
// this special receiver type is available for access inside the function

/*
General syntax
--------------

func (t type) methodName(param type) return type {

}
*/

//----------------------------------------------------
type Person struct {
	fname string
	lname string
	age   int
}

func showDetails(person Person) {
	fmt.Println(person)
}

func (p Person) showDetails() {
	fmt.Println(p)
}

//----------------------------------------------------
type Employee struct {
	fname  string
	lname  string
	salary float64
}

func (e Employee) showDetails() {
	fmt.Println(e)
}

// Value receiver method
func (e Employee) updateName(fname, lname string) {
	e.fname = fname
	e.lname = lname
}

// Pointer receiver method
func (e *Employee) updateNameV2(fname, lname string) {
	e.fname = fname
	e.lname = lname
}

//----------------------------------------------------

type Address struct {
	city  string
	state string
}

type Student struct {
	name  string
	grade string
	Address
}

func (a *Address) update(city, state string) {
	if city != "" {
		a.city = city
	}
	if state != "" {
		a.state = state
	}
}

func (s *Student) showDetails() {
	fmt.Println(s.name, s.grade, s.Address)
}

//----------------------------------------------------
type CustomString string

func (cs *CustomString) concat(s CustomString) {
	*cs += s
}

//----------------------------------------------------

func main() {

	john := Person{"John", "Lane", 10}
	showDetails(john)
	john.showDetails()

	// What is the purpose of methods if functions can do relatively same things?

	// Functions should be used largely as pure functions, whereas
	// Methods extend behaviour programming to types, emulating object oriented behaviour
	// Also, Methods with same names can be declared for different receiver types, functions can't
	// this allows the idea of interfaces to apply in Go, similar to Java

	harry := Employee{"Harry", "Potter", 25000}
	harry.showDetails() // {Harry Potter 25000}
	harry.updateName("Harry", "Gopher")
	harry.showDetails() // {Harry Potter 25000}

	// Why is it still the same?

	// Due to harry employee received as a value in method
	// To pass as reference, we need to pass address of harry

	// pointer receivers
	(&harry).updateNameV2("Harry", "Gopher")
	harry.showDetails() // {Harry Gopher 25000}

	// however, Go allows shorthand, and even for a pointer receiver method
	// explicit passing of address is not required

	harry.updateNameV2("Gopher", "Gopher")
	harry.showDetails() // {Gopher Gopher 25000}

	// similar to Promoted fields in structs
	// Go provides a similar convenience for methods to be called in for nested structs

	// example
	bestStudent := Student{
		"Michael",
		"A+",
		Address{
			"Jersey City",
			"New J",
		},
	}
	bestStudent.showDetails() // Michael A+ {Jersey City New J}
	// update though being a method of Address struct, can be invoked using student object
	// BUT
	// this is only when Address here is an anonymous field in Student struct
	bestStudent.update("", "New Jersey")
	bestStudent.showDetails() // Michael A+ {Jersey City New Jersey}

	// Methods can be created on types other than structs as well, such as int, string etc
	// Rules:
	// 1. Methods can be part of same package as of the type definition package
	// 2. Create alias types and the declare methods on the alias types

	// example
	var greet CustomString = "Hello"
	greet.concat(" World")
	fmt.Println(greet)
}
