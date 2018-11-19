package main

import "fmt"

func main() {
	var words [2]string
	words[0] = "hello"
	words[1] = "world"
	fmt.Println(words[0], words[1])

	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	fmt.Println("length", len(nums))
	fmt.Println("last element", nums[len(nums)-1])
}
