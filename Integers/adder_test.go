package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	actualSum := Add(2, 2)
	expectedSum := 4

	if actualSum != expectedSum {
		t.Errorf("\nActual: %d\nExpected: %d", actualSum, expectedSum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
