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

	// Since both Rectangle and Circle both implement the Shape interface
	// by implementing all the methods in Shape (here area method)
	// Rectangle and Cirle both are Shape type objects as well

	// Example 1
	var s1, s2 Shapes = rect, circle
	fmt.Println(s1.area())
	fmt.Println(s2.area())

	// Example 2
	shapes := []Shapes{rect, circle}
	for _, v := range shapes {
		fmt.Println("Area of", reflect.TypeOf(v), "=", v.area())
	}

	var s3 Shapes
	fmt.Println(reflect.TypeOf(s1)) // main.Rectangle
	fmt.Println(reflect.TypeOf(s2)) // main.Circle
	fmt.Println(reflect.TypeOf(s3)) // <nil>
}
