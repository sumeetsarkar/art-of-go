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

func main() {

	john := Person{"John", "lname", 10}
	showDetails(john)
	john.showDetails()

	// What is the purpose of methods if functions can do relatively same things?

	// Functions should be used largely as pure functions, whereas
	// Methods extend behaviour programming to types, emulating object oriented behaviour
	// Also, Methods with same names can be declared for different receiver types, functions can't
	// this allows the idea of interfaces to apply in Go, similar to Java
}
