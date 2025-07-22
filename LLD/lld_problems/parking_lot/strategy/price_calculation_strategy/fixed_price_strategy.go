package pricecalculationstrategy

import "solid_design/parking_lot/entity"

type FixedPrice struct {
	price float64
}

func (f FixedPrice) CalculatePrice(ticket entity.Ticket) float64 {
	return f.price
}