package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"rectangle": {
			shape: Rectangle{
				Width:  4.0,
				Height: 2.0,
			},
			want: 8.0,
		},
		"circle": {
			shape: Circle{
				Radius: 10.0,
			},
			want: 314.1592653589793,
		},
		"triangle": {
			shape: Triangle{
				Base:   12.0,
				Height: 6.0,
			},
			want: 36.0,
		},
	}

	for name, test := range areaTests {
		t.Run(name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %g, want %g", got, test.want)
			}
		})
	}
}
