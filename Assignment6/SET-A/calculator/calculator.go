package calculator

import "errors"

// Perform calculation based on operator
func Calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}
