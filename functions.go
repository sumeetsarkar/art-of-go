package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
	fmt.Println(swap(1, 2))
}
