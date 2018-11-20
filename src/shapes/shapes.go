package main

import "fmt"
import "shapes/rectangle"
import "shapes/rectangle/rectutils"

func main() {
	fmt.Println("Shapes Main Package")

	length := 10.10
	breadth := 2.4

	area := rectangle.Area(length, breadth)
	perimeter := rectangle.Perimeter(length, breadth)

	fmt.Println("Rectangle Area", area)
	fmt.Println("Rectangle Perimeter", perimeter)

	rectangle.SomeFuncNeedingHelp()

	// rectangle.helper()	// helper not exported, as it starts with lowercase

	rectutils.SomeRectUtils()
}
