package LearnStruct

import "math"

type Shapes interface {
	Perimeter() float64
	Area() float64
}

// 矩形
type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64  {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64  {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64  {
	return 0
}

