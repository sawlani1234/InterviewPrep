package shape


type Recatangle struct {
	length int
	breadth int 
}

func NewRectangle(length,breadth int) Recatangle {
	return Recatangle{length,breadth}
}

func (r Recatangle) ComputeArea() float64 {
	return float64(r.length*r.breadth)
}