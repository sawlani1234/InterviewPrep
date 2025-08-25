package enum

type PieceType string

const (
	INVALID PieceType = "INVALID"
	PAWN    PieceType = "PAWN"
	KING    PieceType = "KING"
	QUEEN   PieceType = "QUEEN"
	KNIGHT  PieceType = "KNIGHT"
	BISHOP  PieceType = "BISHOP"
	ROOK    PieceType = "ROOK"
)
