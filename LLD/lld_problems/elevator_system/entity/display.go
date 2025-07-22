package entity

import (

	"solid_design/elevator_system/enum"
)

type Display struct {

	dir enum.Direction 
	floor int 
	status enum.State

}

func NewDisplay() *Display {
	return &Display{dir: enum.INVALID,floor: 0,status: enum.IDLE}
}

func (d *Display) GetFloor() int {
	return d.floor
}

func (d *Display) GetDirection() enum.Direction {
	return d.dir
}

func (d *Display) GetStatus() enum.State {
	return d.status
}

func (d *Display) Set(dir enum.Direction,floor int,status enum.State) {
	d.dir = dir 
	d.floor = floor
	d.status = status
}

func (d *Display) SetStatusAndDirection(status enum.State,dir enum.Direction) {
	d.status=status
	d.dir = dir
}

