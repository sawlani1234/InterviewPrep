package entities

import "errors"

const maxPercentage = 100 

type PercentageSplitStrategy struct {
	
}


func (p PercentageSplitStrategy) Splits(amount int,splits []Split) ([]Split,error) {

	sum := 0
	newSplits := make([]Split,0)
	for i:=0;i<len(splits);i++ {
		sum += splits[i].GetAmount()
		newSplits = append(newSplits, NewSplit(splits[i].GetUser(),(splits[i].GetAmount()*amount)/100))
	}

	if sum != maxPercentage {
		return []Split{},errors.New("incorrect percentage split")
	}
	return newSplits,nil
}