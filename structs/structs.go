package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (tr Triangle) Area() float64 {
	return tr.Height * tr.Base / 2
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}
