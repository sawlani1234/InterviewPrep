package entity

import (
	"fmt"
	"solid_design/elevator_system/enum"
	"strconv"
)

/**

it will have hashmap and check if the request is needed to be served ?

**/
type LookStrategy struct {
	e Elevator
	floors int
	lookup map[string]bool
}

func NewLookStrategy(e Elevator,floors int) *LookStrategy{

	return &LookStrategy{e :e,floors: floors, lookup : make(map[string]bool)}
}


func (l *LookStrategy) AddNewRequest(floor int, direction enum.Direction) {
	l.lookup[strconv.Itoa(floor)+"#"+string(direction)] = true
}

// should be trigerred by some goroutine and all
/*

elevator in IDLE state , elevator in MOVING state 

*/
func (l *LookStrategy) MoveElevator() {
	
	if len(l.lookup) > 0  { 

		if l.e.display.GetStatus() == enum.IDLE {
			dir := l.e.display.GetDirection()
			if l.e.display.GetFloor() == l.floors {
			    	dir  = enum.DOWN
			} else if l.e.display.GetFloor() == 0 {
				dir = enum.UP
			}
			cur := strconv.Itoa(l.e.display.GetFloor()) + "#" + string(dir)
			if l.lookup[cur] {
				fmt.Println("serving request ",cur)
				delete(l.lookup,cur)
			}
			l.e.display.SetStatusAndDirection(enum.MOVING,dir)
		} else {
			dir := l.e.display.GetDirection()
			
			if l.e.display.GetFloor() == l.floors {
				dir = enum.DOWN 
			} else if l.e.display.GetFloor() == 0 {
				dir = enum.UP
			}
			
			var f int
			if dir == enum.UP {
				f = l.e.display.GetFloor()+1
				l.e.display.Set(dir,f,enum.MOVING)
			} else {
				f = l.e.display.GetFloor()-1
				l.e.display.Set(dir,f,enum.MOVING)
			}

			cur := strconv.Itoa(f)+"#"+string(dir)
			if l.lookup[cur] {
				delete(l.lookup,cur)
				fmt.Println("Proccesed request ",cur)
			    l.e.display.SetStatusAndDirection(enum.IDLE,dir)
			}
			
		}
		return 
	}
	l.e.display.SetStatusAndDirection(enum.IDLE,enum.INVALID)
	
}
