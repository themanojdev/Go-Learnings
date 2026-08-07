package main

import "math"

type  Shape interface {
	Area() float64
	Perimeter() float64
}

type rectangle struct {
	width , height float64
}

func (rect rectangle) Area() float64 {
	return rect.height * rect.width
}

func (rect rectangle) Perimeter() float64 {
	return 2 * (rect.height+rect.width)
}

// when you need to check the interface is implementing the rectangle struct
func check(s Shape) {
	var _ Shape = rectangle{}
	_ = s

	var _ Shape = Cricle{}
}

type Cricle struct {
	radius float64
}

func (c Cricle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Cricle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

}
