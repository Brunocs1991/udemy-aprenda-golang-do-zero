package shapes_test

import (
	"testing"

	. "advanced/shapes"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := Rectangle{Width: 10, Height: 5}
		expected := 50.0
		if r.Area() != expected {
			t.Errorf("Expected %f, got %f", expected, r.Area())
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := Circle{Radius: 5}
		expected := 78.5
		if c.Area() != expected {
			t.Errorf("Expected %f, got %f", expected, c.Area())
		}
	})
}
