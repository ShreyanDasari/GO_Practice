package main

import (
	"fmt"
)

type Shape interface{
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Circle struct{
	radius float64
}
func(c Circle) Area() float64{
	return 3.14 * c.radius*c.radius
}

func main() {
	sq := Square{side: 5}
	ci := Circle{radius: 3}
	shapes := []Shape{sq,ci}
	for _,shape:=range shapes{
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
	}
}
