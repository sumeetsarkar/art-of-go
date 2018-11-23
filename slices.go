package main

import (
	"fmt"
)

// Slices are representation of arrays
// It is a convenient, flexible and powerful wrapper on top of an array.
// Slices do not own any data on their own.
// They are the just references to existing arrays.

func main() {
	arr := [...]int{1, 2, 3}

	// creates a slice of arr from pos 1 to 1
	arrslice := arr[1:2]
	fmt.Println(arrslice)

	// create slice and return immediately
	arrslice = []int{1, 2, 3, 4}
	fmt.Println(arrslice)

	arrslice2 := arrslice
	arrslice2[0] = 4
	fmt.Println(arrslice, arrslice2) // [4 2 3 4] [4 2 3 4]

	colors := [...]string{"R", "G", "B"}
	// create a full length slice of an array
	colorsSlice := colors[:]
	fmt.Println(colors, colorsSlice) // [R G B] [R G B]

	colors[0] = "R1"
	colorsSlice[1] = "G1"
	fmt.Println(colors, colorsSlice) // [R1 G1 B] [R1 G1 B]
}
