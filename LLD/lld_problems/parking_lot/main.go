package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"solid_design/parking_lot/entity"
	"solid_design/parking_lot/enums"
	"solid_design/parking_lot/manager"
	"strconv"
	"sync"
)

const (
	END = "END"
)

var (
	twoWheelerPSM  manager.TwoWheelerParkingSpaceManager
	fourWheelerPSM manager.FourWheelerParkingSpaceManager
	once           sync.Once
	tickets       []entity.Ticket
)

func init() {
	once.Do(func() {
		twoWheelerPSM = manager.NewTwoWheelerParkingSpaceManager()
		fourWheelerPSM = manager.NewFourWheelerParkingSpaceManager()
	})
}


func stringify(a interface{}) string {
	by, _ := json.Marshal(a)
	return string(by)
}

func main() {
	
	printParkingSpace(enums.TwoWheeler)
	fmt.Println("\n------ WELCOME TO SAWLANI PARKING TWO WHEELER AND FOUR WHEELER---\n\n Press 1 for vehicle entry \n Press 2 for vehicle exit \n Type END for ending\n")
		
	scanner := bufio.NewScanner(os.Stdin)
		
	for scanner.Scan() {
		input := scanner.Text()
		if input == END {
			fmt.Println("turning off parking lot program")
			break
		}

		if input == "1" {
			fmt.Println(" ---- Enter Vehicle number")
			scanner.Scan()
			no := scanner.Text()

			fmt.Println(" ---- Enter Vehicle type")
			scanner.Scan()
			ty := scanner.Text()
			vehicleType, _ := strconv.Atoi(ty)

			vehicle := entity.NewVehicle(no, enums.MapIntToVehicleType[vehicleType])
			
			entryGate := entity.NewEntryGate(
				vehicle,
				GetParkingSpaceManager(vehicle.VehicleType))
			
			parkingSpace, err := entryGate.ParkingSpaceManager.FindParkingSpace()
			fmt.Println("found parking space :",stringify(parkingSpace))
			if err != nil {
				fmt.Println(err.Error())
				continue;
			}
			
			err = entryGate.ParkingSpaceManager.BookParkingSpace(parkingSpace)
			if err != nil {
				fmt.Println(err.Error())	
			}
			ticket := entity.NewTicket(len(tickets),vehicle,parkingSpace)
			tickets = append(tickets,ticket)
			fmt.Println("your ticket ", stringify(ticket))
			
			fmt.Println("\n--- print parking space ---\n")
			printParkingSpace(enums.MapIntToVehicleType[vehicleType])
		} else {
			fmt.Println("---- Enter ticket id -- ")
			scanner.Scan()
			ticketId := scanner.Text()
			ticketIDInt,_ := strconv.Atoi(ticketId)
			ticket,err := GetTicketFromTicketID(ticketIDInt)
			if err !=nil {
				fmt.Println(err.Error())
				continue
			}
			exitGate := entity.NewExitGate(ticket, GetParkingSpaceManager(ticket.Vehicle.VehicleType))
			exitGate.ParkingSpaceManager.EmptyParkingSpace(ticket.ParkingSpace)
			money := exitGate.Ticket.ParkingSpace.GetPriceCalculationStrategy().CalculatePrice(ticket)
			fmt.Println("exited vehicle ", stringify(ticket.Vehicle))
			fmt.Println("money to be paid ", money)
			printParkingSpace(enums.TwoWheeler)
		}

	}

}

func GetTicketFromTicketID(id int) (entity.Ticket,error) {
	for _,ticket := range tickets {
		if ticket.Id == id {
			return ticket,nil
		}
	}

	return entity.Ticket{},errors.New("no ticket found")
}


func printParkingSpace(vehicleType enums.VehicleType) {
	switch vehicleType {
	case enums.TwoWheeler:
		for _,ps := range twoWheelerPSM.GetParkingSpaces() {
			fmt.Println(stringify(ps)," Occuppied: " ,ps.IsOccupied())
			
		}
	case enums.FourWheeler:
	default:	
	
	}
}

func GetParkingSpaceManager(vehicleType enums.VehicleType) entity.ParkingSpaceManager {
	switch vehicleType {
	case enums.TwoWheeler:
		return &twoWheelerPSM
	case enums.FourWheeler:
		return &fourWheelerPSM
	default:
		return nil
	}
}