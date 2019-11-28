package perim

import "testing"

func assertEqualFloat(got, want float64, t *testing.T) {
	if got != want {
		t.Errorf("expected %.2f but got %.2f", want, got)
	}
}
func checkArea(want float64, shape Shape, t *testing.T) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("expected %g but got %g", want, got)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assertEqualFloat(got, want, t)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12,Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v expected %g but got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}

type Shape interface {
	Area() float64
}