package golang_united_school_homework

import (
	"errors"
	"fmt"
	//"fmt"
	"reflect"
)

var (
	errorOutOfIndex = errors.New("capacity is out of box's capacity")
	errorExistance  = errors.New("shape doesn't exist")
	errorCircle	    = errors.New("cicrles don't exist")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity < 1 {
		return fmt.Errorf("insufficient capacity")
	}
	
	if len(b.shapes) < b.shapesCapacity { 
		b.shapes = append(b.shapes, shape)
		return nil
	}else {
		return fmt.Errorf("out of index")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (s Shape, e error) {
	if len(b.shapes) != 0 &&  i < len(b.shapes) {
		return b.shapes[i], nil
	}
	return nil, errorExistance
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if len(b.shapes) != 0 &&  i < len(b.shapes) {
		n := NewBox(b.shapesCapacity)
		n.shapes = append(n.shapes, b.shapes[:i]...)
		b.shapes = append(n.shapes, b.shapes[:i+1]...)
		return b.GetByIndex(i)
	}
	return nil, errorExistance
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	b.shapes[i] = shape
	return b.ExtractByIndex(i)
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	cap := b.shapesCapacity
	for i, shape := range b.shapes {
		if reflect.TypeOf(shape).String() == "Circle" {
			b.ExtractByIndex(i)
		}
	}
	if cap == b.shapesCapacity {return errorCircle}
	return nil
}	
