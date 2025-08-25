package cell

import (
	"chess_design/position"
)

type Cell struct {
	Piece    Piece
	Position position.Position
}

func NewCell(piece Piece, position position.Position) *Cell {
	return &Cell{piece, position}
}


