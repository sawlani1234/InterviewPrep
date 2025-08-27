package controller

import (
	"container/list"
	"cricbuzz/design/player"
)

type BowlingController struct {
	bowlers *list.List
}


func NewBowlingController(bowlers []player.Player) *BowlingController{

	l := list.New()

	for _,bowler := range bowlers {
		l.PushBack(bowler)
	}
	return &BowlingController{l}
}

func (b BattingController) getNextBowler() player.Player{
    return player.Player{}
}