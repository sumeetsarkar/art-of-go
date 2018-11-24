package main

import (
	"fmt"
)

// Regular multiplier function
func mul(a, b int) int {
	return a * b
}

// Higher Order Functions
func doubleit(fn func(a, b int) int) func(num int) int {
	wrapper := func(num int) int {
		// Due to property of closures, fn is accessible inside wrapper
		return fn(2, num)
	}
	return wrapper
}

// Higher Order Functions
func times(fn func(a, b int) int) func(multiplier int) func(num int) int {
	wrapper := func(multiplier int) func(num int) int {
		wrapper2 := func(num int) int {
			// Due to property of closures,
			// fn and multiplier is accessible inside wrapper2
			return fn(multiplier, num)
		}
		return wrapper2
	}
	return wrapper
}

func main() {
	// Anonymous Functions
	// Declare the function, just like a type declaration,
	// for functions the signature needs to be mentioned, param list with type and return type
	var greetUser func(name string) string

	// Define the function
	greetUser = func(name string) string {
		return ("Hello " + name)
	}

	// Invoke the function
	result := greetUser("Sumeet")
	fmt.Println(result)

	// or simply use shorthand
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(10, 20))

	twice := doubleit(mul)  // returns wrapper function, which can take num to double it
	fmt.Println(twice(100)) // 200
	fmt.Println(twice(200)) // 400

	times5 := times(mul)(5)
	fmt.Println(times5(100)) // 500
}
