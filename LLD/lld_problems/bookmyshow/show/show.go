package show

import (
	"time"

	"bookmyshow/audi"
	"bookmyshow/enum"
	"bookmyshow/movie"
	"bookmyshow/show_seat"
	pricingstrategy "bookmyshow/strategy/pricing_strategy"
)

type Show struct {
	Id int 
	TimeStart time.Time 
	TimeEnd   time.Time 
	Movie movie.Movie 
	Audi  audi.Audi
	ShowSeats []show_seat.ShowSeat
}

func GetPricingStrategy(seatType enum.SeatType) pricingstrategy.PricingStrategy {
	 if seatType == enum.Gold {
		return pricingstrategy.GoldPricingStrategy{}
	 } else if seatType == enum.Premium {
		return pricingstrategy.ExecutivePricingStrategy{}
	 } else {
		return pricingstrategy.ReclinerPricingStrategy{}
	 }
}

func NewShow(Id int,TimeStart time.Time,TimeEnd time.Time,Movie movie.Movie,
Audi audi.Audi) *Show {

	ShowSeat := []show_seat.ShowSeat{}
	for i:=0;i<len(Audi.Seats);i++ {

		ShowSeat = append(ShowSeat, *show_seat.NewShowSeat(
			Audi.Seats[i],
			GetPricingStrategy(Audi.Seats[i].SeatType),
		))
	}
	return &Show{
		Id,
		TimeStart,
		TimeEnd,
		Movie,
		Audi,
		ShowSeat,
	}
}