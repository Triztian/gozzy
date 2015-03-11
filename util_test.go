package gozzy

import "testing"

func TestLinspace(t *testing.T) {
	a, b := 0.0, 5.0
	n := 10
	expected := []float64{0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0}
	lin := linspace(a, b, n)
	t.Log("Linspace (%f, %f, %v): %v\n", a, b, n, lin)
	t.Log("Expected (%f, %f, %v): %v\n", a, b, n, expected)

	if len(lin) != len(expected) {
		t.Error("Unexpected length: %d", len(lin))
	}

	for i := 0; i < len(lin); i++ {
		if lin[i] != expected[i] {
			t.Error("Unexpected element: %v -> %v, should be %v", i, lin[i], expected[i])
		}
	}
}
