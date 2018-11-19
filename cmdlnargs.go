package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println("number of args passed", len(args))

	for i := range args {
		fmt.Println(args[i])
	}
}
