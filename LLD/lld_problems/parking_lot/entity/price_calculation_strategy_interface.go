package entity

type PriceCalculationStrategy interface {
	CalculatePrice(ticket Ticket) float64
}