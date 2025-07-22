package pricecalculationstrategy

import (
	"time"
	"solid_design/parking_lot/entity"
)


type MinutePrice struct {
}

func (m MinutePrice) CalculatePrice(ticket entity.Ticket) float64 {
	return time.Now().Sub(ticket.GetEntryTime()).Minutes() * 0.2
}