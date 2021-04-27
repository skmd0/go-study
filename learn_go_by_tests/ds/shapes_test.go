package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 6.0}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36.0},
	}
	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.want {
				t.Errorf("got %g want %g", got, tc.want)
			}
		})
	}
}
