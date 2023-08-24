package main

import "fmt"

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func interface_demo() {
	shapes := []Shape{Circle{Radius: 2.0}, Rectangle{Width: 3.0, Height: 4.0}}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
