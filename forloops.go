package main

import "fmt"

func computesum1(arr []int) int {
	sum := 0
	// short hand for loop
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func computesum2(arr []int) int {
	sum := 0
	// traditional for loop
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := computesum1(nums)
	fmt.Println(result)

	result = computesum2(nums)
	fmt.Println(result)

	for {
		break
	}
	fmt.Println("Infinte loop break")
}
