package cell

import (
	"chess_design/enum"
	"chess_design/utils"
	"fmt"
)

type BishopRuleStrategy struct {
}

func (r BishopRuleStrategy) IsMoveRuleAllowed(cells [][]*Cell, start Cell, end Cell) bool {

	colDiff := utils.GetAbs(end.Position.Col - start.Position.Col)
	rowDiff := utils.GetAbs(end.Position.Row - start.Position.Row)

	if end.Piece.GetPieceColor() == start.Piece.GetPieceColor() {
		return false
	}

	if colDiff == rowDiff {
		if start.Position.Row < end.Position.Row   && start.Position.Col < end.Position.Col {
			col := start.Position.Col+1
			for row := start.Position.Row + 1; row < end.Position.Row; row++ {
				if cells[row][col].Piece.pieceType != enum.INVALID {
					fmt.Println("here budyy ", row,col, " : ",cells[row][col])
					return false
				}
				col++
			}
		} else if start.Position.Row > end.Position.Row   && start.Position.Col > end.Position.Col   {
			col := start.Position.Col-1
			for row := start.Position.Row - 1; row > end.Position.Row; row-- {
				if cells[row][col].Piece.pieceType != enum.INVALID {
					return false
				}
				col--
			}
		} else if start.Position.Row < end.Position.Row && end.Position.Col < start.Position.Col {
			col := start.Position.Col-1
			for row := start.Position.Row + 1; row < end.Position.Row; row++ {
				if cells[row][col].Piece.pieceType != enum.INVALID {
					return false
				}
				col--
			}
		} else {
			col := start.Position.Col+1
			for row := start.Position.Row - 1; row > end.Position.Row; row-- {
				if cells[row][col].Piece.pieceType != enum.INVALID {
					fmt.Println("here ",cells[row][col])
					return false
				}
				col++
			}

		}
		return true 
	}
	return false
}
