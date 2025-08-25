package cell

import (
	"chess_design/utils"
	"fmt"
)

type PawnRuleStrategy struct {
}

func (r PawnRuleStrategy) IsMoveRuleAllowed(cells [][]*Cell, start Cell, end Cell) bool {

	rowDiff := utils.GetAbs(end.Position.Row - start.Position.Row)
	colDiff := utils.GetAbs(end.Position.Col - start.Position.Col)

	fmt.Println("here in pawn strategy ",end," : " ,start)

	return (rowDiff == 1) && (colDiff == 0) || ((colDiff == 1) && (rowDiff == 1) && end.Piece.GetPieceColor() != start.Piece.GetPieceColor())
}
