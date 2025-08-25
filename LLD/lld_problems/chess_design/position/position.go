package position


type Position struct {
	Row int 
	Col int 
}


func NewPosition(Row,Col int) Position {
	return Position{Row,Col}
}