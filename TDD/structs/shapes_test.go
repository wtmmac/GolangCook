package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(3.0, 5.0)
	want := 16.00
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		checkArea(t, r, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		shape := Circle{10.0}
		checkArea(t, shape, 314.1592653589793)
	})
}

func TestAreaWithTable(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, v := range areaTests {
		t.Run(v.name, func(t *testing.T) {
			got := v.shape.Area()
			if got != v.want {
				t.Errorf("%#v got %g want %g", v.shape, got, v.want)
			}
		})
	}
}
