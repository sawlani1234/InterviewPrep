package show_seat

import (
	"bookmyshow/enum"
	"bookmyshow/seat"
	"bookmyshow/strategy/pricing_strategy"
)


type ShowSeat struct {
	Seat seat.Seat
	status enum.Status 
	pricingStrategy pricingstrategy.PricingStrategy

}

func NewShowSeat(Seat seat.Seat,pricingStrategy pricingstrategy.PricingStrategy) *ShowSeat {
	return &ShowSeat{
		Seat: Seat,
		status: enum.AVAILABLE,
		pricingStrategy: pricingStrategy,
	
	}
}


func (s *ShowSeat) SetStatus(status enum.Status) {
	s.status = status
}

func (s *ShowSeat) GetStatus() enum.Status {
	return s.status
}

func (s *ShowSeat) GetPriceForSeat() int{
	return s.pricingStrategy.GetPrice()
}