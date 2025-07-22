package lld


type Board struct {
	cells []Cell
}


func NewBoard(cells []Cell) Board {

	return Board{cells}
}

func (b Board) GetSize() int {
	return len(b.cells)
}