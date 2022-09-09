package structs_methods_interfaces

import (
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{length: 10, breadth: 10},
			hasArea: 100.0,
		},
		{
			name:    "Circle",
			shape:   Circle{radius: 10},
			hasArea: 314.1592653589793,
		},
		{
			name:    "Triangle",
			shape:   Triangle{base: 12, height: 6},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("\n%#v\nActual: %g\nExpected: %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
