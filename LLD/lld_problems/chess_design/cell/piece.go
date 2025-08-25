package cell

import (
	"chess_design/enum"
)


type Piece struct {
	pieceColor enum.PieceColor
	pieceType enum.PieceType
	moveRuleStrategy MoveRuleStrategy
}

func NewPiece(pieceColor enum.PieceColor,pieceType enum.PieceType) Piece{
    
	return Piece{
		pieceColor: pieceColor,
		pieceType: pieceType,
		moveRuleStrategy: GetPieceMoveStrategy(pieceType),
	}
}

func (p Piece) GetPieceColor() enum.PieceColor {
	return p.pieceColor
}

func (p Piece) GetPieceMoveStrategy() MoveRuleStrategy{
	return p.moveRuleStrategy
}

func (p Piece) SetPieceType(pieceType enum.PieceType) {
	p.pieceType = pieceType
}


func GetEmptyPiece() Piece {
	return  Piece{pieceType: enum.INVALID}
}

func (p Piece) GetPieceType() enum.PieceType {
	return p.pieceType
}