package main

import "fmt"

var a int = 10
var b int

// a = 10
// syntax error: non-declaration statement outside function body

func main() {
	var local int
	local = 20
	fmt.Println("local", local)

	fmt.Println("a", a)
	fmt.Println("b", b)

	b = 10
	// c = 20	// undefined: c
	d := 10
	// e := 20	// e declared and not used
	fmt.Println("a+b+d", a+b+d)

	var name string = "Sumeet Sarkar"
	fmt.Println("name:", name)

	var (
		abc = 10
		def = 20
	)
	fmt.Println("abc", abc)
	fmt.Println("def", def)
}
