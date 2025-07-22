package shape

import "fmt"

type ShapeFactory struct {
}

func NewShapeFactory() ShapeFactory {
	return ShapeFactory{}
}

func (s ShapeFactory) GetShapeFactory(name string, length, breadth int) (Shape, error) {

	if name == "Square" {
		return NewSquare(length, breadth), nil
	} else if name == "Rectangle" {
		return NewRectangle(length, breadth), nil
	} else if name == "Triangle" {
		return NewTriangle(length, breadth), nil
	} else {
		return nil, fmt.Errorf("invalid shape requested :%v", name)
	}
}
