package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"solid_design/car_rental/entity"
	"solid_design/car_rental/enum"
	"sort"
	"strconv"
	"time"
)


type VehicleRentalSystem struct {
	users []entity.User
	stores []entity.Store
}

func initialise() *VehicleRentalSystem {
	vehicle1 := entity.NewCar("KA-01-2999","Shubham",enum.ACTIVE,20,enum.CAR)
	vehicle2 := entity.NewCar("KA-01-1000","Swarnim",enum.ACTIVE,30,enum.CAR)
	vehicle3 := entity.NewBike("KA-01-2000","Paji",enum.ACTIVE,50,enum.BIKE)
	vehicle4 := entity.NewBike("KA-01-4000","Mental",enum.ACTIVE,620,enum.BIKE)
	
	reservationManager := entity.NewReservationManager()
	
	vehicleManager := entity.NewVehicleManager()
	vehicleManager.Add(vehicle1)
	vehicleManager.Add(vehicle2)
	vehicleManager.Add(vehicle3)
	vehicleManager.Add(vehicle4)
	
	store := entity.NewStore(1,reservationManager,vehicleManager,entity.NewLocation(560037,"Spice Garden,Marathalli"))

	vehicleRentalSystem := NewVehicleRentalSystem()
	
	vehicleRentalSystem.AddUser(entity.NewUser(1,"20202-11-222","Shubham Sawlani"))
	vehicleRentalSystem.AddUser(entity.NewUser(2,"20233-11-333","Swarnim Jaiswal"))
	vehicleRentalSystem.AddStore(store)

	return vehicleRentalSystem

}

func bookingFlow(vehicleRentalSystem *VehicleRentalSystem,scanner *bufio.Scanner) {
	// 1. Enter User id
	fmt.Println("Enter user id") 
	scanner.Scan()
	input := scanner.Text()
	
	userId,err := strconv.Atoi(input)
	printError(err)
	
	user,err := vehicleRentalSystem.GetUser(userId)
	printError(err)
	
	// 2. Enter Pincode
	fmt.Println("Enter your pincode")
	scanner.Scan()
	input = scanner.Text()

	pincode,err := strconv.Atoi(input)
	printError(err)

	// 3. Get Store for that pincode , later can change to multiple stores for that pincode
	store,err := vehicleRentalSystem.GetStore(pincode)
	printError(err)

	// 4. Get Vehicles for that pincode
	vehicles := store.GetVehicleManager().GetVehicle(enum.CAR)

	fmt.Println("Here are the vehicles for the selected stor ,select one id ")

	for i:=0;i<len(vehicles);i++ {
		fmt.Println(stringify(vehicles[i]))
	}

	scanner.Scan()
	input = scanner.Text()

	// 5. Choose a vehicle
	id,err := strconv.Atoi(input)
	printError(err)

	fmt.Println("Reserving selected vehicle id ",id)
	vehicle,err  := store.GetVehicleManager().GetVehicleWithVehicId(id)

	printError(err)

	vehicle.SetStatus(enum.BOOKED)
	store.GetVehicleManager().Update(vehicle)

	layout := "02:01:2006"
	
	// 6. Enter from date
	fmt.Println("eneter booking from date")
	scanner.Scan()
	input = scanner.Text()
	
	fromDate,err := time.Parse(layout,input)
	printError(err)

	// 7. Enter to date
	fmt.Println("eneter booking to date")
	scanner.Scan()
	input = scanner.Text()
	
	toDate,err := time.Parse(layout,input)
	printError(err)

	
	reservation := entity.NewReservation(user,vehicle,enum.SCHEDULED,fromDate,toDate)
	store.GetReservationManager().Add(*reservation)

	// 8. Compute Amount
	amount := reservation.Bill.GenerateBillAmount()

	fmt.Println("you have to pay amount",amount)

	//9. Enter method to pay
	fmt.Println("choose method to pay \n 1: Credit Card \n 2: UPI \n 3: Debit Card")
	scanner.Scan()
	input = scanner.Text()
	id,err = strconv.Atoi(input)
	
	printError(err)

	strategy,err := entity.GetPaymentStrategy(id)
	printError(err)
	
	err = strategy.Pay()
	printError(err)

	reservation.SetStatus(enum.SUBMITTED)
	err = store.GetReservationManager().Update(*reservation)
	printError(err)

}


// TODO: incomplete implementation
func cancelationFlow() {

}

// TODO: incomplete implementation
func completeFlow() {

}

func main() {


	vehicleRentalSystem := initialise()

	fmt.Println("HELLO to SHUBHAM SAWLANI RENTAL SYSTEM")
	
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {

	// first choide is when user wants to book , other operation can be when user wants to cancel a booking and when user wnats to complete the booking

		fmt.Println("Choose \n 1 for booking , \n 2 for Cancellation, \n 3 for Completing")

		scanner.Scan()
		input := scanner.Text()
		choice,err := strconv.Atoi(input)
		printError(err)

		if choice == 1 {
			bookingFlow(vehicleRentalSystem,scanner)

		} else {
			fmt.Println("invalid choide")
			os.Exit(0)
		}
	}
}

func printError(err error) {
	if err != nil {
	fmt.Println("error ",err.Error())
	os.Exit(0)
	}
}

func stringify(a interface{}) string {
	bytes,_ := json.Marshal(a)
	return string(bytes)
}

func NewVehicleRentalSystem() *VehicleRentalSystem{
	return &VehicleRentalSystem{
		users: make([]entity.User, 0),
		stores: make([]entity.Store, 0),
	}
}


func (v *VehicleRentalSystem) AddUser(user entity.User) {
	v.users = append(v.users, user)
}

func (v *VehicleRentalSystem) GetUser(userId int) (entity.User,error) {

	idx := sort.Search(len(v.users),func (i int) bool  {
		return v.users[i].GetId() == userId
	})

	if idx == len(v.users) {
		return entity.User{},errors.New("no user found")
	}

	return v.users[idx],nil

}



func (v *VehicleRentalSystem) AddStore(store entity.Store) {
	v.stores = append(v.stores, store)
}


func (v *VehicleRentalSystem) GetStore(pincode int) (entity.Store,error) {

	idx := sort.Search(len(v.stores),func (i int) bool  {
		return v.stores[i].GetLocation().GetPincode() == pincode
	})

	if idx == len(v.stores) {
		return entity.Store{},errors.New("no store found")
	}

	return v.stores[idx],nil

}