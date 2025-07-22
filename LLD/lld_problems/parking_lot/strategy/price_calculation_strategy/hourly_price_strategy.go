package pricecalculationstrategy

import (
	"time"
	"solid_design/parking_lot/entity"
)

type HourlyPrice struct {
}

func NewHourlyPrice() HourlyPrice {
	return HourlyPrice{}
}

func (h HourlyPrice) CalculatePrice(ticket entity.Ticket) float64 {
	return time.Now().Sub(ticket.GetEntryTime()).Seconds() * 10
}