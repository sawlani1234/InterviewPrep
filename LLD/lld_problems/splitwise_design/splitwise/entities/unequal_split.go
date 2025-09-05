package entities

import "errors"

type UnequalSplitStrategy struct {
}

func NewUnequalSplitStrategy() UnequalSplitStrategy { 
	return UnequalSplitStrategy{}
}

func (e UnequalSplitStrategy) Split(amount int, splits []Split) ([]Split,error) {

	sum := 0
	for i := 0; i < len(splits); i++ {
		sum += splits[i].GetAmount()
	}

	if sum != amount {
		return []Split{}, errors.New("incorrect split")
	}
	return splits,nil
}