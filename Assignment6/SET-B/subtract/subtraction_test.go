package subtraction

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 5},
		{20, 8, 12},
		{15, 15, 0},
		{-5, -10, 5},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
