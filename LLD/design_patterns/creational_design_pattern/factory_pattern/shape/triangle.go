package shape

type Triangle struct {
	length int
	breadth int 
}

func NewTriangle(length,breadth int) Triangle {
	return Triangle{length,breadth}
}

func (t Triangle) ComputeArea() float64 {
	return float64(t.length*t.breadth)*0.5
}