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

	// capacity vs length
	sportsslice := []string{
		"football",
		"cricket",
		"basketball",
		"volleyball",
		"handball",
		"hockey",
		"boxing",
		"wrestling",
		"tennis",
		"badminton",
	}
	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the FIRST element in the slice.
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// capacity: 10 length: 10 [football cricket basketball volleyball handball hockey boxing wrestling tennis badminton]

	sportsslice = sportsslice[:8]
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// capacity: 10 length: 8 [football cricket basketball volleyball handball hockey boxing wrestling]

	sportsslice = sportsslice[:6]
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// capacity: 10 length: 6 [football cricket basketball volleyball handball hockey]

	sportsslice = sportsslice[:4]
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// capacity: 10 length: 4 [football cricket basketball volleyball]

	sportsslice = sportsslice[2:]
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// Note: capacity is counting from the FIRST element in the slice.
	// since, we provide a start index, which is not 0, the count begins from 2 - 10,
	// resulting capacity to be 10 - 2 = 8
	// the slice however was of len 4, now which becomes 2
	// capacity: 8 length: 2 [basketball volleyball]

	sportsslice = sportsslice[2:]
	fmt.Println("capacity:", cap(sportsslice), "length:", len(sportsslice), sportsslice)
	// capacity: 6 length: 0 []
}
