package weight


type LbsWeight interface {
	GetWeightInLbs() int
}



type LbsImpl struct {
   w int 
}

func NewLbsImpl(w int) *LbsImpl{
	return &LbsImpl{w}
}

func (l LbsImpl) GetWeightInLbs() int {
	return l.w
}