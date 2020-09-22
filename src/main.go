package main

import (
	"fmt"
	"math"
)

func main() {
	var r Rectangle = Rectangle{1.5, 6}
	var c Circle = Circle{2.5}

	fmt.Printf("r area = %v", r.area())
	fmt.Printf("c area = %v", c.area())
	fmt.Println(r)
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	r.height = 2
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
