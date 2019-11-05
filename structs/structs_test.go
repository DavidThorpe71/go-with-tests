package structs

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf(`got %g want %g`, got, want)
		}
	}

	t.Run(`rectangles`, func(t *testing.T) {

		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run(`circles`, func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{
			shape:   Rectangle{Width: 12, Height: 6},
			hasArea: 72,
		}, {
			shape:   Circle{Radius: 10},
			hasArea: 314.1592653589793,
		}, {
			shape:   Triangle{Width: 12, Base: 6},
			hasArea: 35.0,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf(`%#v got %.2f want %.2f`, tt.shape, got, tt.hasArea)
		}
	}
}
