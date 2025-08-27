package controller

import "cricbuzz/design/player"


type BattingController struct {
	players []player.Player
	striker player.Player
	nonStriker player.Player
}


func (b *BattingController) getNextPlayer() player.Player {
	return player.Player{}
}


func (b *BattingController) setStriker(p player.Player) player.Player {
   	return p
}

func (b *BattingController) setNonStriker(p player.Player) player.Player {
   	return p
}