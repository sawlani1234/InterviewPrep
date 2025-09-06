package main

import "fmt"

/*
1. You can add up everything as feature  to base object
2. To pervent class explsoion from base class for every feature  -
	- Base pizza , base pizza + Mushroom , base pizza + Cheese , base pizza + etc

*/

type Pizza interface {
	GetPrice() int
}

type VegMania struct {

}

func (v VegMania) GetPrice() int {
	return 100 
}

func NewVegMania() VegMania {
	return VegMania{}
}

type Maraghretta  struct {

}

func (m Maraghretta) GetPrice() int {
	return 120 
}

func NewMarghretta() Maraghretta {
	return Maraghretta{}
}


type ExtraChesseTopping struct {
	p Pizza 
}



func (e ExtraChesseTopping) GetPrice() int {
	return e.p.GetPrice()+20
}

func NewExtraCheeseTopping(p Pizza) ExtraChesseTopping {
	return ExtraChesseTopping{p}
}
type ExtraMayoTopping struct {
	p Pizza
}

func (e ExtraMayoTopping) GetPrice() int {
	return  e.p.GetPrice()+30
}

func NewExtraMayoTopping(p Pizza) ExtraMayoTopping {
	return ExtraMayoTopping{p}
}

func main() {
   fmt.Println(NewExtraMayoTopping(NewExtraCheeseTopping(NewMarghretta())).GetPrice())
}