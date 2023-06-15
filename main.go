package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radus float64
}

func (c Circle) Area() float64 {
	return c.radus * c.radus * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radus
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width * 2) + (r.height * 2)
}

func PrintSaheDetails(s Shape) {
	fmt.Println("Area is ", s.Area())
	fmt.Println("Perimeter is", s.Perimeter())

}

func main() {
	rect := Rectangle{width: 4, height: 5}
	circ := Circle{radus: 7}
	PrintSaheDetails(rect)
	PrintSaheDetails(circ)

}
