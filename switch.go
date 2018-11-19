package main

import (
	"fmt"
)

func redColorString() string {
	fmt.Println("redColorString func called")
	return "red"
}

func blackColorString() string {
	fmt.Println("blackColorString func called")
	return "black"
}

func main() {
	color := "blue"

	switch color {
	case redColorString():
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	case blackColorString():
		fmt.Println("The color is black")
	default:
		fmt.Println("Default... no color selected")
	}
}

/*
Output...

redColorString func called
The color is blue
*/
