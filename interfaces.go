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

	// Empty Interface
	// It has zero methods to implement and hence all types implement it

	// example

	logit := func(data interface{}) {
		fmt.Println(data)
	}
	logit(s1)
	logit(s2)
	logit(s3)

	// Since, interfaces can accept any type,
	// Type assertions is needed to deduce the underlying type
	// This can be done using reflect.TypeOf or else using following

	var data1 interface{} = "string type data"
	var data2 interface{} = 1000

	// fmt.Println(data1.(int)) // panic: interface conversion: interface {} is string, not int
	fmt.Println(data2.(int)) // works, sunce data2 is actually int

	// to do type assertions safer way instead
	// Note: this only works for interface {} type vars

	v, ok := data1.(string)
	fmt.Println(v, ok)
	if ok {
		fmt.Println("data1 is string")
	}

	// Type switch
	// <your-interface-var>.(type) can only be used with switch

	switch data2.(type) {
	case int:
		fmt.Println("data2 is int, value=", data2.(int))
	case string:
		fmt.Println("data2 is string, value=", data2.(string))
	default:
		fmt.Println("unkown type")
	}

	// Example 2 for type assertion

	// create rect1 and circle1 objects
	rect1 := Rectangle{40, 10}
	circle1 := Circle{15}

	// declare 2 vars of type Shapes
	var shape1, shape2 Shapes

	// assign rect1 and circle1 to shape1 and shape2 resp.
	shape1, shape2 = rect1, circle1

	// declare an anonymous function taking Shape type vars
	computeArea := func(i interface{}) {
		// get the actual type
		switch v := i.(type) {
		case Shapes:
			fmt.Println("area of the shape is", v.area())
		case int:
			fmt.Println("v is int, value=", v)
		case string:
			fmt.Println("v is string, value=", v)
		default:
			fmt.Println("unkown type, cannot compute area")
		}
	}

	computeArea(shape1)
	computeArea(rect1) // since rect1 implemente Shape
	computeArea(shape2)
	computeArea(circle1) // since circle1 implements Shape

	computeArea(data1)
	computeArea(data2)

	// not allowed
	// v := data1.(type)	// use of .(type) outside type switch
	// fmt.Println(v)

	person1 := Person{"Sumeet", "Sarkar"}
	// var md Metadata = person1	// Person does not implement Metadata (info method has pointer receiver)
	var md Metadata = &person1
	md.info()

	// However for methods the relaxation of interchangeably using pointer vs values to call methods is applied
	person1.pointerReceiver()
	(&person1).pointerReceiver()
	person1.valueReceiver()
	(&person1).valueReceiver()

	// But as per concept of value and point receivers, only the method with pointer receiver changes the data
	person1.info() // Sumeet 2 Sarkar
}

type Metadata interface {
	info()
}

type Person struct {
	fname, lname string
}

func (p *Person) info() {
	fmt.Println(p.fname, p.lname)
}

func (p *Person) pointerReceiver() {
	p.fname = "Sumeet 2"
	fmt.Println("called from pointer receiver")
}

func (p Person) valueReceiver() {
	p.lname = "Sarkar 2"
	fmt.Println("called from value receiver")
}
