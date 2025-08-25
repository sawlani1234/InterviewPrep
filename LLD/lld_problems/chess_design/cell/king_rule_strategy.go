package cell

import "chess_design/utils"

type KingRuleStrategy struct {
}

func (r KingRuleStrategy) IsMoveRuleAllowed(cells [][]*Cell, start Cell, end Cell) bool {

	colDiff := utils.GetAbs(end.Position.Col - start.Position.Col)
	rowDiff := utils.GetAbs(end.Position.Row - start.Position.Row)

	if colDiff == 0 && rowDiff == 0 {
		return false
	}

	return colDiff <= 1 && rowDiff <= 1 && end.Piece.GetPieceColor() != start.Piece.GetPieceColor()
}
