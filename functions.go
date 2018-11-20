package main

import (
	"fmt"
)

/*
Basic syntax of functions
-------------------------
func funcname(param type) returntype {
}
*/

func fireandforget() {
	fmt.Println("fire and forget")
}

// single return value
func add(x, y int) int {
	return x + y
}

// multiple return values
func swap(x, y int) (int, int) {
	return y, x
}

// named return values
func rect(l, b float64) (area, perimeter float64) {
	perimeter = 2 * (l + b)
	area = l * b
	return
}

func main() {
	fireandforget()

	fmt.Println("add", add(10, 20))

	x, y := 1, 2
	fmt.Println("x, y", x, y)

	x, y = swap(x, y)

	fmt.Println("swapped x, y")
	fmt.Println("x, y", x, y)

	val1, val2 := rect(10, 20)
	fmt.Println(val1, val2)

	// suppress recieve of one of the params using,
	// Blank Identifier
	area, _ := rect(10, 20)
	fmt.Println(area)
}
