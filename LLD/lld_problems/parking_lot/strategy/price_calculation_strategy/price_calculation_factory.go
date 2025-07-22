package pricecalculationstrategy

import (
	"solid_design/parking_lot/entity"
)


func NewPriceCalculationStrategy(priceType string) entity.PriceCalculationStrategy {
	switch priceType {
	case "fixed":
		return FixedPrice{price: 100}
	case "hourly":
		return HourlyPrice{}
	case "minute":
		return MinutePrice{}
	default:
		return nil
	}
}
