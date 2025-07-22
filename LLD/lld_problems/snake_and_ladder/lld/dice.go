package lld

import "math/rand"

type Dice struct {
	noOfDice int
}

func NewDice(noOfDice int) Dice {
	return Dice{noOfDice}
}

func (d Dice) RollDice() int {

	sum := 0
	for i := 1; i <= d.noOfDice; i++ {
		sum+=rand.Intn(7)
	}
	return sum
}