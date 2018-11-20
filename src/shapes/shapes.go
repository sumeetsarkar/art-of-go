package main

import (
	"fmt"
	"shapes/rectangle"
	"shapes/rectangle/rectutils"
)

func init() {
	fmt.Println(">>>>>>>>>> shapes main initialized")
}

func main() {
	fmt.Println("!!! Shapes Main Package starts")

	length := 10.10
	breadth := 2.4

	area := rectangle.Area(length, breadth)
	perimeter := rectangle.Perimeter(length, breadth)

	fmt.Println("Rectangle Area", area)
	fmt.Println("Rectangle Perimeter", perimeter)

	rectangle.SomeFuncNeedingHelp()

	// rectangle.helper()	// helper not exported, as it starts with lowercase

	rectutils.SomeRectUtils()

	fmt.Println("!!! Shapes Main Package ends")
}

/*
Output

>>>>>>>>>> rectuils utils initialized
>>>>>>>>>> rectangle recthelper initialized
>>>>>>>>>> rectangle rectprops initialized
>>>>>>>>>> shapes main initialized
!!! Shapes Main Package starts
Rectangle Area 24.24
Rectangle Perimeter 25
Some Func Needing Help...
Some Rect Utils...
rectangle Helper
Some Rect Utils...
!!! Shapes Main Package ends
*/
