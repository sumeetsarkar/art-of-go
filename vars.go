package main

import "fmt"

var a int = 10
var b int

// a = 10
// syntax error: non-declaration statement outside function body

func main() {
	// variable declaration
	var local int
	// variable assignment
	local = 20
	fmt.Println("local", local)

	// variable declaration and assignment together
	var name string = "Sumeet Sarkar"
	fmt.Println("name:", name)

	// multiple variable declaration and assignment
	// 2 variables declared, 2 variable assigments expected
	var fname, lname string = "sumeet", "sarkar"
	fmt.Println(fname, lname)

	// var firstName, lastName string = "sumeet"	// assignment mismatch: 2 variables but 1 values

	// printing variables defined in global
	fmt.Println("a", a)
	fmt.Println("b", b)

	// variable assignment of global
	b = 10
	// c = 20	// undefined: c, trying to assign a varaiable not declared yet

	// ---------------------------------------
	// implicit type inference
	// requires no type specification, infers type from value assigned
	var e = 10
	// var f	// expecting type

	// multiple variable declaration and assignment
	var g, h = 20, 30

	// --------------------------------------
	// shorthand declaration
	// implict type inference and declaration done together
	d := 10

	// multiple variable declaration and assignment
	x, y := 20, 30

	// e := 20	// e declared and not used

	fmt.Println("a+b+d+e+g+h+x+y", a+b+d+e+g+h+x+y)

	// ---------------------------------------
	// special syntax to declare multiple variables of different types
	var (
		abc       = 10
		firstName = "sumeet"
		lastName  = "sarkar"
		above18   = true
	)
	fmt.Println("abc", abc)
	fmt.Println("firstName", firstName)
	fmt.Println("lastName", lastName)
	fmt.Println("above18", above18)
}
