package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	wanted := 40.0

	if got != wanted {
		t.Errorf("\nActual: %.2f\nExpected: %.2f", got, wanted)
	}
}

func TestArea(t *testing.T) {

	t.Run("should return area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		wanted := 100.0

		if got != wanted {
			t.Errorf("\nActual: %g\nExpected: %g", got, wanted)
		}
	})

	t.Run("should return area of circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		wanted := 314.1592653589793

		if got != wanted {
			t.Errorf("\nActual: %g\nExpected: %g", got, wanted)
		}
	})

}
