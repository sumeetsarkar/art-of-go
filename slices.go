package main

import (
	"fmt"
)

// slices are representation of arrays

func main() {
	arr := []int{1, 2, 3}

	arrslice := arr[1:2] // creates a slice of arr from pos 1 to 1
	fmt.Println(arrslice)

	// create slice and return immediately
	arrslice = []int{1, 2, 3, 4}
	fmt.Println(arrslice)

	arrslice2 := arrslice
	arrslice2[0] = 4
	fmt.Println(arrslice2)
}
