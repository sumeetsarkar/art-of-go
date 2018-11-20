package rectangle

import (
	"fmt"
	utils "shapes/rectangle/rectutils"
)

func init() {
	fmt.Println(">>>>>>>>>> rectangle rectprops initialized")
}

func Area(l, b float64) float64 {
	return l * b
}

func Perimeter(l, b float64) float64 {
	return 2 * (l + b)
}

func SomeFuncNeedingHelp() {
	fmt.Println("Some Func Needing Help...")
	utils.SomeRectUtils()
	helper()
}
