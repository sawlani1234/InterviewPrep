package main

import (
	"bookmyshow/audi"
	"bookmyshow/booking_service"
	"bookmyshow/controller"
	"bookmyshow/enum"
	"bookmyshow/movie"
	"bookmyshow/seat"
	"bookmyshow/show"
	"bookmyshow/theatre"
	"time"
)


var theatreController *controller.TheatreController
var MovieController *controller.MovieController
var ShowController *controller.ShowController


func initialise() {

	MovieController = controller.NewMovieController()
	theatreController = controller.NewTheatreController()
	ShowController = controller.NewShowController()
	// Initialise Movie controller
	bahubali := movie.NewMovie(0,"Bahubhali","Action/Drama")
	sonOfSardar2 := movie.NewMovie(1,"Son of Sardar 2","Comedy")
	MovieController.AddMovie(enum.Bengaluru,*bahubali)
	MovieController.AddMovie(enum.Bengaluru,*sonOfSardar2)
	MovieController.AddMovie(enum.Delhi,*sonOfSardar2)

	// Initialise Theatre Controller 
	seats := make([]seat.Seat,10)

	for i:=0;i<len(seats);i++ {
		seats[i].Id = i 
		seats[i].Number = i
		if i < 3 {
 		seats[i].SeatType = enum.Gold
		} else if i >= 4 && i<=8 {
		 seats[i].SeatType = enum.Premium	
		} else {
			seats[i].SeatType = enum.Recliner
		}
	}

	audi1 := audi.NewAudi(0,"A1",seats)
	theatre1 := theatre.NewTheatre(0,"Shantiniketan",[]audi.Audi{*audi1})

	seats2 := make([]seat.Seat,10)
	
	for i:=0;i<len(seats2);i++ {
		seats2[i].Id = i 
		seats2[i].Number = i
		if i < 3 {
 		seats2[i].SeatType = enum.Gold
		} else if i >= 4 && i<=8 {
		 seats2[i].SeatType = enum.Premium	
		} else {
			seats2[i].SeatType = enum.Recliner
		}
	}
	audi2 := audi.NewAudi(1,"A2",seats2)
	theatre2 := theatre.NewTheatre(1,"PVR Delhi ke",[]audi.Audi{*audi2})

	theatreController.AddTheatre(enum.Bengaluru,*theatre1)
	theatreController.AddTheatre(enum.Delhi,*theatre2)
	

	// Iniitialse ShowsController for movies  

	show1 := show.NewShow(
		0,
		time.Date(2025,1,1,12,0,30,0,time.UTC),
		time.Date(2025,1,1,2,0,30,0,time.UTC),
		*bahubali,
		*audi1,
	)

	show2 := show.NewShow(1,
		time.Date(2025,1,1,3,0,30,0,time.UTC),
		time.Date(2025,1,1,5,0,30,0,time.UTC),
		*sonOfSardar2,
		*audi1,
	)

	show3 := show.NewShow(2,
		time.Date(2025,1,1,3,0,30,0,time.UTC),
		time.Date(2025,1,1,5,0,30,0,time.UTC),
		*sonOfSardar2,
		*audi2,
	)

	ShowController.AddShows(audi1.Id,*show1)
	ShowController.AddShows(audi1.Id,*show2)
	ShowController.AddShows(audi2.Id,*show3)
}






func main() {

	initialise()


	booking_service.
	NewBookingService(theatreController,MovieController,ShowController).
	ProcessBooking()
	

}