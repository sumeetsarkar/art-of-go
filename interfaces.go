package main

import (
	"fmt"
	"math"
	"reflect"
)

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Shapes interface {
	area() float64
}

func main() {

	rect := Rectangle{10, 20}
	circle := Circle{10}

	shapes := []Shapes{rect, circle}
	for _, v := range shapes {
		fmt.Println("Area of", reflect.TypeOf(v), "=", v.area())
	}
}
