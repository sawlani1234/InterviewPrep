package lld


type Player struct {
	id int 
	name string
	currentPosition int 
}


func NewPlayer(id int,name string) *Player {

	return &Player{
		id:id,
		name:name,
		currentPosition:0,
	}

}


func (p *Player) GetId() int {
	return  p.id
}


func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPosition() int {
	return p.currentPosition
}

func (p *Player) SetPosiition(n int)  {
	p.currentPosition = n
}