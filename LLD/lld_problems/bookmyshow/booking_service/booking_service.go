package booking_service

import (
	"bookmyshow/controller"
	"bookmyshow/enum"
	"bookmyshow/user"
	"fmt"
)


type BookingService struct {
    theatreController *controller.TheatreController
	movieController *controller.MovieController
	showController *controller.ShowController
}

func NewBookingService( theatreController *controller.TheatreController, movieController *controller.MovieController,
showController *controller.ShowController) *BookingService {
	return  &BookingService{
		theatreController,
		movieController,
		showController,
	}
}


func (b *BookingService) ProcessBooking() {

	for ;; {
	fmt.Println("Type C to continue/start booking or E to exist")
	var ch string
	fmt.Scanf("%s",&ch)

	if ch == "E" {
		fmt.Println("Thanks for using booking sir exiting now ")
		break
	}
	fmt.Println("Type your Name sir ")
	var s string 
	fmt.Scanf("%s",&s)
	user := user.NewUser(s)

	fmt.Println("Select Id city Sir")

	cities := enum.GetAllCities()
	fmt.Println("Id  :  City")
	for i:=0;i<len(cities);i++ {
		fmt.Println(i, " : " ,cities[i])
	}

	var cityID int

	fmt.Scanf("%d",&cityID)

	movies := b.movieController.MovieCityMap[cities[cityID]]

	fmt.Println("\nSelect from  Available movies")
	
	for i:=0;i<len(movies);i++ {
		fmt.Println(i, " : " ,movies[i].Name, " : ",movies[i].Genre)
	}

	var movieId int 

	fmt.Scanf("%d",&movieId)

	fmt.Printf("All Theatre and there shows for the movie :%s\n",movies[movieId].Name)

	fmt.Println("Select Show ID and Audi ID")
	
	theatres := b.theatreController.CityToTheatreMap[cities[cityID]]
	
	for i:=0;i<len(theatres);i++ {
		audis := theatres[i].Audis

		for j:=0;j<len(audis);j++ {
			shows := b.showController.AudiToShowMap[audis[j].Id]
			
			for k:=0;k<len(shows);k++ {
				if shows[k].Movie.Name == movies[movieId].Name {
					fmt.Println("\n Show ID: ", k,"\n Theatre :",theatres[i].Name, "\n Audi :",audis[j].Name,"\n Audi Id ",audis[j].Id, "\n Start time ", shows[k].TimeStart,  "\n End time ",shows[k].TimeEnd)
				}
		    }
		}
	}

	var showId,audiId int 

	fmt.Scanf("%d %d",&showId,&audiId)

	fmt.Println("Select from available  seats")

	showSeats := b.showController.AudiToShowMap[audiId][showId].ShowSeats

	for i:=0;i<len(showSeats);i++ {
		if showSeats[i].GetStatus() == enum.AVAILABLE {
			fmt.Println(showSeats[i].Seat.Number , " : ", showSeats[i].GetPriceForSeat())
		}
	}

	var seatNumber,price int 
	fmt.Scanf("%d",&seatNumber)

	for i:=0;i<len(showSeats);i++ {
		if showSeats[i].Seat.Number == seatNumber {
			showSeats[i].SetStatus(enum.BOOKED)
			price = showSeats[i].GetPriceForSeat()
			break
		}
	}

	// Get Theatre

	var theatreId int 
	for i:=0;i<len(theatres);i++ {
		for j:=0;j<len(theatres[i].Audis);j++ {
			if theatres[i].Audis[j].Id == audiId {
				theatreId = i
				break
			}
		}
	}

	fmt.Println("Booking complete here are the details")
	fmt.Println("User :",user.Name)
	fmt.Print("Theatre :",theatres[theatreId].Name)
	fmt.Println("Movie :",movies[movieId].Name)
	fmt.Println("Audi :",audiId)
	fmt.Println("Price :", price)

	fmt.Println("Available Seats in Theatre ")

	showSeats2 := b.showController.AudiToShowMap[audiId][showId].ShowSeats
	for i:=0;i<len(showSeats2);i++ {
		if showSeats2[i].GetStatus() == enum.AVAILABLE {
			fmt.Println(showSeats2[i].Seat.Number)
		}
	}

}

}