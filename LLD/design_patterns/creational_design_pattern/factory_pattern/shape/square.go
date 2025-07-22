package shape

import "fmt"

type Square struct {
	length int
}

func NewSquare(length,breadth int) Square {

	if length != breadth {
		fmt.Println("not a valid square")
	}
	return Square{length}
}

func (s Square) ComputeArea() float64 {
	return float64(s.length*s.length)
}