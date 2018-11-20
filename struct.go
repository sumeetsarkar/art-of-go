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

	// anonymous structs
	sumeet := struct {
		name string
		age  int
	}{
		name: "Sumeet Sarkar",
		age:  28,
	}
	fmt.Println(sumeet)

	// anonymous fields in structs
	type AnonyFields struct {
		string
		int
		bool
	}
	magic := AnonyFields{
		string: "magic",
		int:    100,
		bool:   true,
	}
	fmt.Println(magic)
	fmt.Println(magic.string)
	magic.int += 100
	fmt.Println(magic.int)

	// anonymous struct with anonymous fields
	nextlevelmagic := struct {
		string
		int
		bool
		// bool	// not allowed, duplicate field bool
	}{"nextlevelmagic", 100, true}
	fmt.Println(nextlevelmagic)
	fmt.Println(nextlevelmagic.bool)
}
