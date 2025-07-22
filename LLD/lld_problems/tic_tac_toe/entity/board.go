package entity

import (
	"errors"
	"fmt"
	"solid_design/tic_tac_toe/common"
	"solid_design/tic_tac_toe/enum"
)

type Board struct {
	board [][]enum.Piece
}

func NewBoard(n int) Board{

	board := make([][]enum.Piece,n)

	for i:=0;i<n;i++ {
		board[i] = make([]enum.Piece, n)
	}

	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			board[i][j] = enum.W
		}
	}

	return Board{board}

}

func (b *Board) IsValidRowColumn(i,j int) bool {
	n := len(b.board)
	
	return i >=0 && i < n && j>=0 && j<n
}

func (b *Board) IsEmpty(i,j int) bool {
	return b.IsValidRowColumn(i,j) && b.board[i][j] == enum.W
}


func (b *Board) CheckWinnerTopRow() bool {
	
	n := len(b.board)
	first := b.board[0][0]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[0][i] != first {
			isWinner = false
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerBottomRow() bool {
	
	n := len(b.board)
	first := b.board[n-1][0]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[n-1][i] != first {
			isWinner = false
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerLeftColumn() bool {
	
	n := len(b.board)
	first := b.board[0][0]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[i][0] != first {
			isWinner = false
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerRightColumn() bool {
	
	n := len(b.board)
	first := b.board[0][n-1]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[i][n-1] != first {
			isWinner = false
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerRightDiagonal() bool {
	
	n := len(b.board)
	first := b.board[0][n-1]

	if first == enum.W {
		return false
	}

	isWinner := true 
	
	for i:=1;i<n;i++ {
		if b.board[i][n-i-1] != first {
			isWinner = false 
			break
		}
	}
	return isWinner
}


func (b *Board) CheckWinnerLeftDiagonal() bool {
	
	n := len(b.board)
	first := b.board[0][0]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[i][i] != first {
			isWinner = false 
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerMiddleVertical() bool {
	n := len(b.board)
	first := b.board[0][n/2]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[i][n/2] != first {
			isWinner = false 
			break
		}
	}
	return isWinner
}

func (b *Board) CheckWinnerMiddleHorizontal() bool {
	n := len(b.board)
	first := b.board[n/2][0]
	isWinner := true 

	if first == enum.W {
		return false
	}
	
	for i:=1;i<n;i++ {
		if b.board[n/2][i] != first {
			isWinner = false 
			break
		}
	}
	return isWinner
}

func (b *Board) GetWinnerPiece() (bool,string,enum.Piece) {
	isWinner := b.CheckWinnerTopRow()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[0][0]
	}

	isWinner = b.CheckWinnerLeftColumn()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[0][0]
	}

	isWinner = b.CheckWinnerBottomRow()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[len(b.board)-1][0]
	}

	isWinner = b.CheckWinnerRightColumn()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[0][len(b.board)-1]
	}

	isWinner = b.CheckWinnerLeftDiagonal()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[0][0]
	}

	isWinner = b.CheckWinnerRightDiagonal()
	
	if isWinner {
		return true,common.EMPTY_STRING,b.board[0][len(b.board)-1]
	}

	isWinner = b.CheckWinnerMiddleHorizontal() || b.CheckWinnerMiddleVertical()

	if isWinner {
		return true ,common.EMPTY_STRING,b.board[len(b.board)/2][len(b.board)/2]
	}


	if b.IsAnyEmpty() {
		return false,common.NO_RESULT,enum.W 
	}

	return false,common.TIE,enum.W
}

func (b *Board) IsAnyEmpty() bool {
	n := len(b.board)
	
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			if b.board[i][j] == enum.W {
				return true 
			}
		}
	}

	return false
}


func (b *Board) Set(i int,j int,piece enum.Piece) error {
	if !b.IsValidRowColumn(i,j) || !b.IsEmpty(i,j) {
		return errors.New("invalid column choosed sir")
	}

	b.board[i][j] = piece
	return nil
}


func (b *Board) PrintBoard() {

	n := len(b.board)
	
	for i:=0;i<n;i++ {
		for j:=0;j<n-1;j++ {
			fmt.Print(string(b.board[i][j]), " | ")
		}
		fmt.Println(string(b.board[i][n-1]))
		if i!=(n-1) {
		fmt.Println("----------")
		}
	}

}