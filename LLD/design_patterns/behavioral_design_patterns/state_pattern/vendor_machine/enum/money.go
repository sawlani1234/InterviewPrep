package enum


type Money int  

const (
	Twenty Money  = 20
	Hundred Money = 100
	FiveHundred Money = 500
)


func (m Money) Int() int {
	return int(m)
}