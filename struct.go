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

	// nested structs
	type Address struct {
		street string
		city   string
	}
	type Employee struct {
		fname   string
		lname   string
		age     int
		address Address
	}
	sam := Employee{
		"Sam",
		"Williams",
		28,
		Address{
			"32 street",
			"SFO",
		},
	}
	fmt.Println(sam)
	fmt.Println(sam.fname, sam.lname, sam.age)
	fmt.Println(sam.address.street, sam.address.city)

	// Promoted fields
	type Employee2 struct {
		fname string
		lname string
		age   int
		Address
	}
	jimmy := Employee2{
		"Jimmy",
		"Jones",
		31,
		Address{
			"54 Avenue",
			"LAX",
		},
	}
	// jimmy gets access to street field from jimmy.address
	// but only if Employee2 does not have a street field
	fmt.Println(jimmy.street, jimmy.city)

	// struct comparision
	person1 := Person{
		"A",
		"B",
		10,
	}
	person2 := Person{
		"A",
		"B",
		10,
	}
	fmt.Println(person1 == person2) // true
	person2.age += 1
	fmt.Println(person1 == person2) // false

	// pointers to a struct
	personPointer := &Person{
		"ABC",
		"DEF",
		30,
	}
	fmt.Println(personPointer)  // &{ABC DEF 30}
	fmt.Println(*personPointer) // {ABC DEF 30}

	// accessing fields in struct using pointer to struct
	fmt.Println((*personPointer).fname) // ABC
	// go provides shorthand accessing for the fields, like regular struct refs
	fmt.Println(personPointer.fname) // ABC

	(*personPointer).fname = "XYZ1"
	fmt.Println(personPointer.fname) // XYZ1
	personPointer.fname = "XYZ2"
	fmt.Println(personPointer.fname) // XYZ2
}
