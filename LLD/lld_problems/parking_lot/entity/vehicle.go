package entity

import "solid_design/parking_lot/enums"

type Vehicle struct {
	RegistrationNumber string
	VehicleType        enums.VehicleType
	id                 int
}

func NewVehicle(registrationNumber string, vehicleType enums.VehicleType) Vehicle {
	return Vehicle{
		RegistrationNumber: registrationNumber,
		id:                 1,
		VehicleType:        vehicleType,
	}
}
