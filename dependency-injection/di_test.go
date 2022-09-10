package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ashutosh")

	got := buffer.String()
	wanted := "Hello, Ashutosh"

	if got != wanted {
		t.Errorf("\nActual: %q\nExpected: %q", got, wanted)
	}
}
