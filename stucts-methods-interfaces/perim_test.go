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
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		checkArea(got, rectangle, t)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		checkArea(got, circle, t)
	})
}

type Shape interface {
	Area() float64
}