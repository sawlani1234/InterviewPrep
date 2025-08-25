package cell

import (
	"chess_design/enum"
	"chess_design/utils"
)

type RookRukeStrategy struct {
}

func (r RookRukeStrategy) IsMoveRuleAllowed(cells [][]*Cell, start Cell, end Cell) bool {

	colDiff := utils.GetAbs(end.Position.Col - start.Position.Col)
	rowDiff := utils.GetAbs(end.Position.Row - start.Position.Row)

	if end.Piece.GetPieceColor() == start.Piece.GetPieceColor() {
		return false
	}

	if colDiff != 0 && rowDiff != 0 {
		return false
	}

	if colDiff == 0 {
		if end.Position.Row > start.Position.Row {
			for row := start.Position.Row + 1; row < end.Position.Row; row++ {
				if cells[row][end.Position.Col].Piece.pieceType != enum.INVALID {
					return false
				}
			}
		} else {
			for row := start.Position.Row - 1; row > end.Position.Row; row-- {
				if cells[row][end.Position.Col].Piece.pieceType != enum.INVALID {
					return false
				}
			}
		}
	}

	if rowDiff == 0 {
		if end.Position.Col > start.Position.Col {
			for col := start.Position.Col + 1; col < end.Position.Col; col++ {
				if cells[end.Position.Row][col].Piece.pieceType != enum.INVALID {
					return false
				}
			}
		} else {
			for col := start.Position.Col - 1; col > end.Position.Col; col-- {
				if cells[end.Position.Row][col].Piece.pieceType != enum.INVALID {
					return false
				}
			}
		}
	}

	return true

}
