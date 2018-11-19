package main

import (
	"fmt"
	"math"
)

func shortStmt(m, n, lim float64) float64 {
	if v := math.Pow(m, n); v < lim {
		return v
	}
	return lim
}

func standardIfElse(age int) string {
	if age < 18 {
		return "Below 18 years old"
	} else if age >= 60 {
		return "Senior citizen"
	} else {
		return "Age group 18 - 59"
	}
}

func main() {
	age := 20
	fmt.Println(standardIfElse(age))

	fmt.Println(shortStmt(2, 3, 10))
	fmt.Println(shortStmt(2, 4, 10))
}
