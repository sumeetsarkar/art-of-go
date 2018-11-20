package main

import "fmt"

func main() {
	a := 10
	var p *int = &a

	fmt.Println("value of a", a)
	fmt.Println("pointer address", p)
	fmt.Println("value from pointer", *p)

	*p += 1
	fmt.Println("value from pointer", *p)

	p = nil
	// *p += 1 // SIGSEGV: Runtime error invalid memory address or nil pointer dereferenc

	var q *int = &a
	fmt.Println("value from pointer", *q)
}
