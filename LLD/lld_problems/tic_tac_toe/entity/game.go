package entity

import (
	"solid_design/tic_tac_toe/common"
	"solid_design/tic_tac_toe/enum"
)


type Game struct {
	currentPlayerTurn int 
    players []Player
	board Board
}

func NewGame(players []Player,board Board) Game {
	return Game{0,players,board}
}


func (g *Game) GetCurrentPlayerTurn() int {
	return g.currentPlayerTurn
}

func (g *Game) GetPlayerforPiece(piece enum.Piece) Player {

	for _,val := range g.players {
		if val.piece == piece {
			return val
		}
	}
	return Player{}
}

func (g *Game) GetWinner() string {

	isWinner,tie,piece := g.board.GetWinnerPiece()

	//fmt.Println(isWinner," : ",tie, " : ",string(piece))

	if isWinner {
		return g.GetPlayerforPiece(piece).name
	}

	if tie == common.TIE {
		return common.TIE
	}

	return common.NO_RESULT
}

func (g *Game) Play(i,j int) error {
     err := g.board.Set(i,j,g.players[g.currentPlayerTurn].piece)
	 if err != nil {
		return err
	 }
	 g.currentPlayerTurn = (g.currentPlayerTurn+1)%len(g.players)
	 return nil
}