package main

import "fmt"

/*
Basic types in go
-----------------

boolean type: bool

string type: string

numeric types:

unsigned int:
	uint8	:	8 bits unsigned integer
	uint16	:	16 bits unsigned integer
	uint32	:	32 bits unsigned integer
	uint64	:	64 bits unsigned integer
	uint	:	32 or 64 bit unsigned interger depending on underlying platform

signed int
	int8	: 8 bit signed integer
	int16	: 16 bit signed integer
	int32	: 32 bit signed integer
	int64	: 64 bit signed integer
	int		: 32 or 64 bit signed integer depending on underlying platform

float32	: 32 byte signed float
float64	: 64 byte signed float (default float64 is used)

byte : 8 bits, alias for uint8 which is unsigned 8 bit integer

rune : alias for int32, signed 32 bit

complex64 : float32 real and float32 imaginary part
complex128 : float64 real and float64 imaginary part

*/

func main() {

	above18 := true
	isTrue := 10 < 20

	fmt.Println(above18, isTrue)

	// type of a variable printed using %T format specifier in Printf method
	num1 := 10
	num2 := 10.10
	name := "Sumeet Sarkar"

	var b byte = 255
	// b = 256	// constant 256 overflows byte

	var r rune = 2147483647
	// r = 2147483648 // constant 2147483648 overflows rune
	fmt.Printf("types: num1 %T, num2 %T, isTrue %T, name %T, b %T, r%T\n", num1, num2, isTrue, name, b, r)

	// type conversions
	// Go needs explicity type conversions for operations between different types

	// sum := num1 + num2	// invalid operation: num1 + num2 (mismatched types int and float64)

	sum1 := float64(num1) + num2
	fmt.Println("sum1 is float", sum1)

	sum2 := num1 + int(num2)
	fmt.Println("sum2 is int", sum2)

	fraction := 10.2 / 2
	fmt.Println(fraction)

	// var numerator = 10.2
	// var denominator = 2
	// fmt.Println(numerator / denominator) // (mismatched types float64 and int)

	const numerator = 10.2
	const denominator = 2
	fmt.Println(numerator / denominator)
}
