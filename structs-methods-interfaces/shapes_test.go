package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	wanted := 40.0

	if got != wanted {
		t.Errorf("\nActual: %.2f\nExpected: %.2f", got, wanted)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	wanted := 100.0

	if got != wanted {
		t.Errorf("\nActual: %.2f\nExpected: %.2f", got, wanted)
	}
}
