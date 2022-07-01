package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcArea() float64 {
	var h float64 = (t.Side*math.Sqrt(3))/2
	
	return 1.0/2.0 * h * t.Side
}

func (t Triangle) CalcPerimeter() float64 {
	return 3*t.Side
}