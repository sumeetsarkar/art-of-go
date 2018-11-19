package main

import "fmt"

const AGE_LIMIT = 18

func main() {
	const num int = 10
	const RED = "RED"

	fmt.Println("num", num)
	fmt.Println("RED", RED)

	// RED = "RED"	// cannot assign to RED
}
