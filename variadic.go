package main

import (
	"fmt"
)

// variadic functions can accept variable number of parameters

func addItems(arr []int, items ...int) []int {
	for _, v := range items {
		// Note: append returns a new slice with appended item,
		// If capacity of the slice cannot accomodate new items to be appended
		// the new slice capacity is doubled

		// read in detail here
		// https://stackoverflow.com/questions/23531737/how-the-slice-is-enlarged-by-append-is-the-capacity-always-doubled
		arr = append(arr, v)
	}
	return arr
}

func change(word ...string) {
	word[0] = "Hi"
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1 2 3 4 5]

	// passing variable number of integers to addItems
	result := addItems(arr, 6)
	result = addItems(result, 7, 8)
	result = addItems(result, 9, 10, 11)
	fmt.Println(result) // [1 2 3 4 5 6 7 8 9 10 11]

	// passing a slice, directly is not allowed for the variadic arguments
	// result = addItems(result, []int{12, 13, 14, 15})

	// syntactical sugar ... to achieve the same
	result = addItems(result, []int{12, 13, 14, 15}...)
	fmt.Println(result)

	// declare greet string array
	greet := [...]string{"Hello", "World"}
	// slice of greet
	greetSlice1 := greet[:]

	change(greetSlice1[0], greetSlice1[1])
	fmt.Println(greetSlice1) // [Hello World]

	change(greetSlice1...)
	fmt.Println(greetSlice1) // [Hi World]

	// Because, if ... is used the greetSlice1 here will be itself be sent
	// to function change, instead of a new slice being created in function
}
