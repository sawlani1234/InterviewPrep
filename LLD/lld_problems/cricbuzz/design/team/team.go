package team

import "cricbuzz/design/player"

type Team struct {
	name string 
	players []player.Player

}


func NewTeam(name string,players []player.Player) *Team {
	return &Team{name,players}
}