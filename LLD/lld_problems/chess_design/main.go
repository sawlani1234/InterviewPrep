package main

import (
	"chess_design/game"
	"chess_design/player"
)

func main() {
	playerA := player.NewPlayer("Shubham")
	playerB := player.NewPlayer("Nishi")
	game.NewGame(playerA, playerB).StartGame()

}
