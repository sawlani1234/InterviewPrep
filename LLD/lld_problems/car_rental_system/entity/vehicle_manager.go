package entity

import (
	"fmt"
	"sort"

	"solid_design/car_rental/enum"
)

type VehicleManager struct {
	vehicles []Vehicle
}


func NewVehicleManager() *VehicleManager {
	return &VehicleManager{vehicles: make([]Vehicle,0)}
}

func (v *VehicleManager) Add(vehicle Vehicle) {
	vehicle.SetId(len(v.vehicles))
	v.vehicles = append(v.vehicles, vehicle)
}

func (v *VehicleManager) Remove(vehicle Vehicle) error {
	idx := -1
	
	for i:=0;i<len(v.vehicles);i++ {
		if v.vehicles[i].GetId() == vehicle.GetId() {
			idx = i 
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("no such vehicle present with id:%v",vehicle.GetId())

	}
	return nil 
}

func (v *VehicleManager) Update(vehicle Vehicle) error {
	for i:=0;i<len(v.vehicles);i++ {
		if v.vehicles[i].GetId() == vehicle.GetId() {
			v.vehicles[i] = vehicle
			return nil 
		}
	}
	return fmt.Errorf("no such vehicle present with id :%v",vehicle.GetId())
}

func (v *VehicleManager) GetVehicle(vehicleType enum.VechileType)[]Vehicle{

	vehicles := []Vehicle{}
	for i:=0;i<len(v.vehicles);i++ {
		if v.vehicles[i].GetVehicleType() == vehicleType  && v.vehicles[i].GetStatus()  == enum.ACTIVE {
			vehicles = append(vehicles, v.vehicles[i])
		}
	}

	return vehicles
}

func (v *VehicleManager) GetVehicleWithVehicId(id int) (Vehicle,error){
	idx := sort.Search(len(v.vehicles),func (i int) bool {
		return v.vehicles[i].GetId() == id
	})

	if idx == len(v.vehicles) {
		return nil,fmt.Errorf("invalid vehicle")
	}
	return v.vehicles[idx],nil
}