package main

import "fmt"

const AGE_LIMIT = 18

func main() {
	const num int = 10
	const RED = "RED"

	fmt.Println("num", num)
	fmt.Println("RED", RED)

	// RED = "RED"	// cannot assign to RED

	// const num2 = math.Pow(2, 8) // const initializer math.Pow(2, 8) is not a constant
}
