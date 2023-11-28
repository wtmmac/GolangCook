package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(3.0, 5.0)
	want := 16.00
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
