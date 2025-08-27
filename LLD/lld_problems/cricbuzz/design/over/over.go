package over

import "cricbuzz/design/ball"


type Over struct {
	overNumber int 
	balls []ball.Ball
}

func NewOver() *Over{
	return &Over{0,make([]ball.Ball, 6)}
}


func (o *Over) StartOver() {

}