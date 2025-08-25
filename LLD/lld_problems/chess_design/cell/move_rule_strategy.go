package cell

import "chess_design/enum"


type MoveRuleStrategy interface {
	IsMoveRuleAllowed(cells [][]*Cell,start Cell,end Cell) bool 
}


func GetPieceMoveStrategy(pieceType enum.PieceType) MoveRuleStrategy {
	switch pieceType {
	case enum.BISHOP :
		return BishopRuleStrategy{}
	case enum.QUEEN:
		return  QueenRuleStrategy{}
	case enum.KING:
		return  KingRuleStrategy{}
	case enum.ROOK :
		return RookRukeStrategy{}
	case enum.PAWN :
		return PawnRuleStrategy{}	
	case enum.KNIGHT :
		return KingRuleStrategy{}	
	}

	return  nil
    
}