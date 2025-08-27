package player

import "cricbuzz/design/enums"

type Player struct {
	name string 
	playerType enums.PlayerType
}


func NewPlayer(name string,playerType enums.PlayerType) *Player {
	return &Player{name,playerType}
}

