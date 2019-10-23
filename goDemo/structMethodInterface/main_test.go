package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{12, 6}, 72.0},
		{&Circle{10}, 314.1592653589793},
		{&Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f wang %.2f", tt.shape, got, tt.want)
		}
	}

}
