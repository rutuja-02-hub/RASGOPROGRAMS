package calculator

import "testing"

func TestAdd(t *testing.T) {
	// Test cases
	tests := []struct {
		a, b     int
		expected int
	}{
		{3, 5, 8},
		{10, 15, 25},
		{-2, -3, -5},
		{0, 5, 5},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
