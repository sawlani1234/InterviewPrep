package entity

import (
	"solid_design/tic_tac_toe/enum"
	"github.com/google/uuid"
)

type Player struct {
	id string 
	name string 
	piece enum.Piece
}

func NewPlayer(name string,piece enum.Piece) Player {
	return Player{
		id : uuid.New().String(),
		name : name,
		piece: piece,
	}
}

func (p Player) GetName() string {
	return p.name
}