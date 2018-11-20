package main

import "fmt"

func main() {
	// declare a map type var of string keys and int values
	var students map[string]int
	// create a map using the map function
	students = make(map[string]int)

	// cannot use make(map[string]string) (type map[string]string) as type map[string]int in assignment
	// students = make(map[string]string) not allowed

	// push items to map
	students["john"] = 98
	students["james"] = 86

	// print map
	fmt.Println(students)

	// get items from map
	fmt.Println("john score", students["john"])

	// key not present in map
	fmt.Println("no one score", students["no one"]) // returns default value of the type, for int 0

	// implicit declaration/ creation of maps using shorthand
	above18 := make(map[string]bool)
	above18["john"] = true
	above18["james"] = false

	fmt.Println("john above 18", above18["john"])

	// delete from map
	delete(above18, "john")
	fmt.Println("john above 18", above18["john"]) // returns default value of the type, for bool false

	// check for presence of key in map
	value, itemPresent := above18["james"]
	// hence check with johnPresent
	if itemPresent {
		fmt.Println("james above18?", value)
	} else {
		fmt.Println("james not found")
	}

	// declaration of map and assigning values at same time
	salaries := map[string]int{
		"optimus":   120000,
		"bumblebee": 100000,
	}
	fmt.Println(salaries)

	// length of map
	fmt.Println("total number of salaries", len(salaries))

	// maps cannot be compared to, other than nil
	// if salaries == above18 {
	// }
	// invalid operation: salaries == above18 (mismatched types map[string]int and map[string]bool)

	var colorcodes map[string]int

	// colorcodes["r"] = 0 // code failure, exit status 2

	if colorcodes == nil {
		colorcodes = make(map[string]int)
	}

	colorcodes["r"] = 0
	fmt.Println("color code for r", colorcodes["r"])
}
