package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	age   int
}

func main() {

	// explicit var declaration
	var john Person
	// assignment with Person struct
	john = Person{
		fname: "John",
		lname: "Mathew",
		age:   29,
	}
	fmt.Println(john)

	// shorthand
	james := Person{
		fname: "James",
		lname: "Bond",
	}
	fmt.Println(james) // age is 0, default of resp type

	// creating stuct without using field names, order to be maintained
	ruby := Person{"Ruby", "Roy", 20} // all fields necessary, in same order
	fmt.Println(ruby)

	var noOne Person
	fmt.Println(noOne) // empty value structure

	// accessing struct values
	var jessica Person
	// field assignment
	jessica.fname = "Jessica"
	jessica.lname = "Parker"
	jessica.age = 28
	// accessing field values
	fmt.Printf("%s %s is %d years old\n", jessica.fname, jessica.lname, jessica.age)
}
