package cell

/*
 
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .
 . . . . . . . .

8*8 chess board
*/
type QueenRuleStrategy struct {
}

func (r QueenRuleStrategy) IsMoveRuleAllowed(cells [][]*Cell,start Cell,end Cell) bool {
	   
	return RookRukeStrategy{}.IsMoveRuleAllowed(cells,start,end) && BishopRuleStrategy{}.IsMoveRuleAllowed(cells,start,end)
}