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

	// iterating arrays

	// method 1, for loop using length
	for i := 0; i < len(rgb); i++ {
		fmt.Println(rgb[i])
	}

	// method 2, for loop using range
	// range returns index and value at the index, for each iteration
	for i, v := range rgb {
		fmt.Printf("arr[%d]: %s\n", i, v)
	}

	// suppress index
	for _, v := range rgb {
		fmt.Println(v)
	}

	// multidimensional array
	arrMulti := [2][2]string{
		{},
	}
	fmt.Println(arrMulti) // [[0 0] [0 0]]

	// assign array in first dimension
	arrMulti[0] = [...]string{"some", "value"}
	arrMulti[1] = [...]string{"another", "value"}
	fmt.Println(arrMulti) // [[some value] [another value]]

	// assign values specific to position in 2d array
	arrMulti[0][1] = "value modified"
	fmt.Println(arrMulti) // [[some value modified] [another value]]

	// Arrays though easy to work with it restricts on its size
}
