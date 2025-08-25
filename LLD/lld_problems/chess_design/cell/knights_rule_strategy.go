package cell

import "chess_design/utils"

type KnightsRuleStrategy struct {
}

func (r KnightsRuleStrategy) checkRowMove(start Cell, end Cell) bool {

	colCheck := utils.GetAbs(end.Position.Col-start.Position.Col) == 1
	rowCheck := utils.GetAbs(end.Position.Row-start.Position.Row) == 2
	return rowCheck && colCheck && start.Piece.GetPieceColor() != end.Piece.GetPieceColor()
}

func (r KnightsRuleStrategy) checkColMove(start Cell, end Cell) bool {
	colCheck := utils.GetAbs(end.Position.Col-start.Position.Col) == 2
	rowCheck := utils.GetAbs(end.Position.Row-start.Position.Row) == 1
	return rowCheck && colCheck && start.Piece.GetPieceColor() != end.Piece.GetPieceColor()
}

func (r KnightsRuleStrategy) IsMoveRuleAllowed(cells [][]*Cell, start Cell, end Cell) bool {

	return r.checkRowMove(start, end) || r.checkColMove(start, end)
}
