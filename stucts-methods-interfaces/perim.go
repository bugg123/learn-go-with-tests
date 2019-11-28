package perim

import "math"


//Rectangle represents the width and heigh of a rectangle
type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle represents the radius of a circle
type Circle struct{
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

//Perimeter used to get the perimeter of rectangle
func Perimeter(rect Rectangle) (perim float64) {
	perim = 2 * (rect.Width + rect.Height)
	return
}

//Area used to get area of a square
func Area(rect Rectangle) (area float64) {
	area = rect.Width * rect.Height
	return
}
