package main

import "fmt"

func main() {
	// declaration of array

	// array of int elements of size 5
	var nums [5]int
	// all elements are defaulted based on type, int is 0, bool is false etc
	nums[0] = 1
	nums[3] = 2
	fmt.Println(nums)

	// shorthand declaration of array
	colors := [...]string{"R", "G", "B"} // compiler infers the length of array based on the elements provided

	// length of array using len()
	fmt.Println(len(colors), colors)

	// type declaration in Array is type + size of the array
	// hence, [10]int and [5]int are not same and cannot be assigned

	arr1 := [10]int{1, 2, 3} // array of length 10, but provided 3 values, rest are zero values
	arr2 := [3]int{1, 2, 3}
	// arr1 = arr2 // not allowed, cannot use arr2 (type [5]int) as type [10]int in assignment
	fmt.Println(arr1, arr2)
	// fmt.Println(arr1 == arr2) // not allowed, invalid operation: arr1 == arr2 (mismatched types [10]int and [5]int)

	// arrays are value types and not referenced
	rgb := [...]string{"R", "G", "B"}
	rgbCopy := rgb
	rgbCopy[0] = "C"
	rgbCopy[1] = "M"
	rgbCopy[2] = "Y"
	fmt.Println(rgb, rgbCopy) // [R G B] [C M Y]
}
