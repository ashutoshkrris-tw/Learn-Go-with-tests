package structs_methods_interfaces

import (
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, wanted float64) {
		t.Helper()
		got := shape.Area()

		if got != wanted {
			t.Errorf("\nActual: %g\nExpected: %g", got, wanted)
		}
	}

	t.Run("should return area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("should return area of circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
