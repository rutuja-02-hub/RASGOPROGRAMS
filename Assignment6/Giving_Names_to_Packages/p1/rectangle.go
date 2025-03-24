package geometry

// Rectangle struct with length and width
type Rectangle struct {
	Length, Width float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Perimeter calculates the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}
