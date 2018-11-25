package main

import (
	"fmt"
)

// Embedding iterfaces

type Engine interface {
	start()
	stop()
}

type Wheel interface {
	drive()
}

type Vehicle interface {
	Engine
	Wheel
}

type Car struct {
	name string
}

func (c Car) start() {
	fmt.Println(c.name, "engine starts")
}

func (c Car) stop() {
	fmt.Println(c.name, "engine stops")
}

func (c Car) drive() {
	fmt.Println(c.name, "drives")
}

func main() {
	var car Vehicle = Car{
		name: "BMWW",
	}
	car.start()
	car.drive()
	car.stop()
}
