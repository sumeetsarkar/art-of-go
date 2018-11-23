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

	// appending to slices
	// not allowed with array
	// append(colors, "C")	// first argument to append must be slice; have [3]string
	newslice := append(colorsSlice, "C")
	newslice = append(newslice, "M", "Y")

	fmt.Println(colorsSlice, newslice)

	// Memory Optimisation
	// Slices hold a reference to the underlying array.
	// As long as the slice is in memory, the array cannot be garbage collected.
	// This might be of concern when it comes to memory management.

	// declare an array of countires
	countries := [...]string{
		"India",
		"USA",
		"UK",
		"Russia",
		"Canada",
		"Australia",
		"China",
		"Japan",
	}
	// create slice of top 3 countires from countries array
	sliceofcountires := countries[:len(countries)-5]
	// declare an array of same size as the slice
	top3arr := make([]string, len(sliceofcountires))
	// copy the contents from the slice to the newly declaraed array
	copy(top3arr, sliceofcountires)
	fmt.Println(top3arr) // [India USA UK]
}
