package lld


type Cell struct {
	number int 
	jump *Jump
}

func NewCell(number int,jump *Jump) Cell {
	return Cell{number,jump}

}

func (c *Cell) GetNumber() int {
	return  c.number
}

func (c *Cell) GetJump() *Jump {
	return c.jump
}