package lld

import (
	"container/list"
	"fmt"
)

type GameManager struct {
	Board   Board
	players *list.List
	dice Dice
}


func NewGameManager(board Board,players *list.List,dice Dice) GameManager {
	return  GameManager{board,players,dice}
}

func (g *GameManager) StartGame() {

	for ;!g.IsWinner(); {
		player := g.getNextPlayer()
		numDice := g.dice.RollDice()
		player.SetPosiition(g.computeNewCell(player,numDice))
		fmt.Println("Player ", player.GetName(),"dice no ", numDice," moved to ",player.GetPosition())
	}

	winner := g.GetWinner()

	fmt.Println("Yahoo , congrats ", winner.name, " you deserved it buddy")
	
}


func (g *GameManager) getNextPlayer() *Player {
	player := g.players.Front().Value.(*Player)
	g.players.Remove(g.players.Front())
	g.players.PushBack(player)
	return player
}

func (g *GameManager) IsWinner() bool {
	return g.GetWinner().GetPosition() == (g.Board.GetSize()-1)
}

func (g *GameManager) GetWinner() *Player {
	return  g.players.Back().Value.(*Player)
}

func (g *GameManager) computeNewCell(player *Player,diceRoll int) int {
	if player.currentPosition+diceRoll >= g.Board.GetSize() {
		return player.currentPosition
	}
	newCell := player.currentPosition+diceRoll
	if g.Board.cells[newCell].jump.GetStart() == -1 { 
		return  newCell
	}
	if g.Board.cells[newCell].jump.GetStart() == newCell {
		return g.Board.cells[newCell].jump.GetEnd()
	}

	return newCell
}