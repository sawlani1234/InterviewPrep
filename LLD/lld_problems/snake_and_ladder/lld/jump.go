package lld


type Jump struct {
	start int
	end int 
}

func NewJump(start,end int) *Jump {
	return &Jump{start,end}
}

func (j *Jump) GetStart() int {
	return j.start
}

func (j *Jump) GetEnd() int {
	return j.end
}

func (j *Jump) SetStart(n int) {
	j.start = n
}

func (j *Jump) SetEnd(n int) {
	j.end = n
}