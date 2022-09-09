package structs_methods_interfaces

type Rectangle struct {
	length  float64
	breadth float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.breadth
}
