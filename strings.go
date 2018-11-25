package main

import (
	"fmt"
	"unicode/utf8"
)

// In go, string is a slice of bytes.
// strings are created by enclosing content within ""

func change(value string) {
	value = "value is changed"
}

func changeByPointer(value *string) {
	*value = "value is changed"
}

func main() {
	// string creation
	greet := "Hello world"
	fmt.Println(greet) // Output: Hello world

	// bytes in strings
	for i := range greet {
		fmt.Printf("%x ", greet[i]) // Output: 48 65 6c 6c 6f 20 77 6f 72 6c 64
		// outputs unicode utf-8 encoded values
	}
	fmt.Println()

	for i := range greet {
		fmt.Printf("%x %c\n", greet[i], greet[i])
	}
	// Output
	// 48 H
	// 65 e
	// 6c l
	// 6c l
	// 6f o
	// 20
	// 77 w
	// 6f o
	// 72 r
	// 6c l
	// 64 d

	//----------------------

	// byte or rune?
	// let us try to try printing bytes and recreate string from those bytes
	// for a string which contains characters outside standard english alphabets

	// Eg: window in hindi is खिड़की
	hindiword := "खिड़की"

	fmt.Println(hindiword) // outputs: खिड़की

	for i := range hindiword {
		fmt.Printf("%c", hindiword[i])
	}
	fmt.Println()

	// Outputs खिड़की as àààààà Why?

	// Note: utf-8 character size is 4 bytes long
	// whereas we were printing 1 byte a time

	// rune is alias for int32, which is 4 bytes long signed int
	// hence, rune represents a unicode character in go

	// create a rune slice
	runeslice := []rune(hindiword)
	for i := range runeslice {
		fmt.Printf("%c", runeslice[i])
	}
	// outputs: खिड़की

	fmt.Println()

	// however, in go, to iterate over a string
	// a rune slice is not needed to be created each time

	// instead, range returns the second argument as a rune
	// so for loop of type as below works fine
	// for _, v := range yourstring {}

	for _, runeValue := range hindiword {
		fmt.Printf("%c", runeValue)
	}
	// outputs: खिड़की

	fmt.Println()

	// now to see how many bytes is occupied by each char in hindiword खिड़की
	for i := range hindiword {
		fmt.Printf("%d ", i)
	}
	// outputs: 3 6 9 12 15
	// which makes it clear that each character includes 3 bytes

	//----------------------
	fmt.Println()

	// length of strings
	fmt.Println("length of hello world ", len(greet), utf8.RuneCountInString(greet))
	fmt.Println("length of खिड़की ", len(hindiword), utf8.RuneCountInString(hindiword))

	// outputs lenght and rune count of strings
	// length of hello world  11 11
	// length of खिड़की  18 6

	//----------------------
	fmt.Println()

	// Strings are immutable

	colorOrange := "Orange"
	fmt.Println(colorOrange) // Outputs: Orange
	// colorOrange[0] = 'R' // cannot assign to colorOrange[0]

	// instead create a rune slice of orange
	colorOrangeSlice := []rune(colorOrange)
	colorOrangeSlice[0] = 'R'

	fmt.Println(colorOrange, string(colorOrangeSlice)) // Outputs: Orange Rrange

	// string concatenation
	greet += " test"
	fmt.Println(greet) // output: Hello world test

	result := append([]rune(greet), '1', '2')
	fmt.Println(string(result)) // output: Hello world test12

	greetnew := "greet before change"
	change(greetnew)
	fmt.Println(greetnew)

	changeByPointer(&greetnew)
	fmt.Println(greetnew)
}
