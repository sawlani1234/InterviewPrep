package board

import (
	"chess_design/cell"
	"chess_design/enum"
	"chess_design/player"
	"chess_design/position"
	"fmt"
	"github.com/fatih/color"
)


type Board struct {
	cells [][]*cell.Cell
	playerToPieceMap map[player.Player][]cell.Piece
}

func NewBoard(playerToPieceMap map[player.Player][]cell.Piece) Board {

	cells := make([][]*cell.Cell,8)

	for i:=0;i<8;i++ {
		cells[i] = make([]*cell.Cell, 8)
	}

	for i:=0;i<8;i++ {
		for j:=0;j<8;j++ {
			cells[i][j] = cell.NewCell(cell.GetEmptyPiece(),position.NewPosition(i,j))
		}
	}

	for j:=0;j<8;j++ {
		pawn := cell.NewPiece(enum.WHITE,enum.PAWN)
		pos := position.NewPosition(1,j)
		cells[1][j] = cell.NewCell(pawn,pos)
	}

	for j:=0;j<8;j++ {
		pawn := cell.NewPiece(enum.BLACK,enum.PAWN)
		pos := position.NewPosition(6,j)
		cells[6][j] = cell.NewCell(pawn,pos)
	}


	// rook
	rook := cell.NewPiece(enum.WHITE,enum.ROOK)
	cells[0][0] = cell.NewCell(rook,position.NewPosition(0,0))
	cells[0][7] = cell.NewCell(rook,position.NewPosition(0,7))

	rook = cell.NewPiece(enum.BLACK,enum.ROOK)
	cells[7][0] = cell.NewCell(rook,position.NewPosition(7,0))
	cells[7][7] = cell.NewCell(rook,position.NewPosition(7,7))

	// knight
	knight := cell.NewPiece(enum.WHITE,enum.KNIGHT)
	cells[0][1] = cell.NewCell(knight,position.NewPosition(0,1))
	cells[0][6] = cell.NewCell(knight,position.NewPosition(0,6))

	knight = cell.NewPiece(enum.BLACK,enum.KNIGHT)
	cells[7][1] = cell.NewCell(knight,position.NewPosition(7,1))
	cells[7][6] = cell.NewCell(knight,position.NewPosition(7,6))

	// bishop
	bishop := cell.NewPiece(enum.WHITE,enum.BISHOP)
	cells[0][2] = cell.NewCell(bishop,position.NewPosition(0,2))
	cells[0][5] = cell.NewCell(bishop,position.NewPosition(0,5))

	bishop = cell.NewPiece(enum.BLACK,enum.BISHOP)
	cells[7][2] = cell.NewCell(bishop,position.NewPosition(7,2))
	cells[7][5] = cell.NewCell(bishop,position.NewPosition(7,5))


	// king 
	king := cell.NewPiece(enum.WHITE,enum.KING)
	cells[0][4] = cell.NewCell(king,position.NewPosition(0,4))

	king = cell.NewPiece(enum.BLACK,enum.KING)
	cells[7][3] = cell.NewCell(king,position.NewPosition(7,3))
	
	// queen
	queen := cell.NewPiece(enum.WHITE,enum.QUEEN)
	cells[0][3] = cell.NewCell(queen,position.NewPosition(0,3))
	
	queen = cell.NewPiece(enum.BLACK,enum.QUEEN)
	cells[7][4] = cell.NewCell(queen,position.NewPosition(7,4))

	
    
	return  Board{cells : cells,playerToPieceMap: playerToPieceMap}

}


func (b Board) CanMove(start cell.Cell,end cell.Cell) bool {
      return start.Piece.GetPieceMoveStrategy().IsMoveRuleAllowed(b.cells,start,end)
}





func (b Board) GetNoOfPieces(color enum.PieceColor) int {
	cnt := 0
	for row := 0;row < len(b.cells);row++ {
		for col := 0;col < len(b.cells[0]);col++ {
			if b.cells[row][col].Piece.GetPieceColor() == color {
				cnt++ 
			}
		}
	}

	return cnt
}



func (b Board) GetCells() [][]*cell.Cell {
	return b.cells
}

func (b Board) MovePiece(start,end cell.Cell) {

	b.cells[end.Position.Row][end.Position.Col].Piece = b.cells[start.Position.Row][start.Position.Col].Piece
	b.cells[start.Position.Row][start.Position.Col].Piece  = cell.GetEmptyPiece() 
	
}

func (b Board) PrintBoard() {

	blue := color.New(color.FgBlue).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()

	fmt.Println("    0  1  2  3  4  5  6  7  8")
	fmt.Println("-------------------------------")
	for i:=0;i<len(b.cells);i++ {
		fmt.Print(i, " | ")
		for j:=0;j<len(b.cells[0]);j++ {
			
			if b.cells[i][j].Piece.GetPieceColor() == enum.WHITE {
			
			if b.cells[i][j].Piece.GetPieceType() == enum.INVALID {
				fmt.Print(" . ")
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.KING {
				fmt.Print(blue(" K "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.QUEEN {
				fmt.Print(blue(" Q "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.BISHOP {
				fmt.Print(blue(" B "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.KNIGHT {
				fmt.Print(blue(" KN "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.PAWN {
				fmt.Print(blue(" P "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.ROOK {
				fmt.Print(blue(" R "))
			}
		} else {

			if b.cells[i][j].Piece.GetPieceType() == enum.INVALID {
				fmt.Print(" . ")
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.KING {
				fmt.Print(green(" K "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.QUEEN {
				fmt.Print(green(" Q "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.BISHOP {
				fmt.Print(green(" B "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.KNIGHT {
				fmt.Print(green(" KN "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.PAWN {
				fmt.Print(green(" P "))
			}
			if b.cells[i][j].Piece.GetPieceType() == enum.ROOK {
				fmt.Print(green(" R "))
			}
		}
	}
	fmt.Println()
	}

	fmt.Println("----------------------------------")
}