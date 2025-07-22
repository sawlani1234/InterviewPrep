package entity

import "solid_design/car_rental/enum"

type VehicleBase struct {

	id int
	vehicleNo string 
	ownerName string 
	status enum.VehicleStatus
	chargePerDay int 
	vechicleType enum.VechileType
}

func NewVehicleBase(id int,vehicleNo string,ownerName string,status enum.VehicleStatus,chargePerDay int,vehicleType enum.VechileType) VehicleBase {
	return VehicleBase{id,vehicleNo,ownerName,status,chargePerDay,vehicleType}
}


type Vehicle interface {
	GetId() int 
	SetId(id int) 
	GetVehicleNo() string 
	GetOwnerName() string 
	GetVehicleType() enum.VechileType
	GetStatus() enum.VehicleStatus 
	GetChargesPerDay() int 
	SetStatus(enum.VehicleStatus) 
}

type Car struct {
	v VehicleBase
}

func NewCar(vehicleNo string,ownerName string,status enum.VehicleStatus,chargePerDay int,vehicleType enum.VechileType) *Car {
	return &Car{v: NewVehicleBase(0,vehicleNo ,ownerName ,status,chargePerDay,vehicleType)}
}

func (c *Car) GetId() int {
   return c.v.id
}

func (c *Car) SetId(id int)  {
	c.v.id = id
}

func (c *Car) GetOwnerName() string {
	return c.v.ownerName
}

func (c *Car) GetVehicleNo() string {
	return c.v.vehicleNo
}

func (c *Car) GetStatus() enum.VehicleStatus {
	return c.v.status
}

func (c *Car) GetChargesPerDay() int {
	return c.v.chargePerDay
}

func (c *Car) SetStatus(status enum.VehicleStatus)  {
	c.v.status = status
}

func (c *Car) GetVehicleType() enum.VechileType {
	return c.v.vechicleType
}



type Bike struct {
	v VehicleBase
}

func NewBike(vehicleNo string,ownerName string,status enum.VehicleStatus,chargePerDay int,vehicleType enum.VechileType)*Bike {
	return &Bike{v: NewVehicleBase(0 ,vehicleNo ,ownerName ,status,chargePerDay,vehicleType)}
}

func (b *Bike) GetId() int {
   return b.v.id
}

func (b *Bike) SetId(id int) {
	 b.v.id = id 
 }

func (b *Bike) GetVehicleNo() string {
	return b.v.vehicleNo
}

func (b *Bike) GetStatus() enum.VehicleStatus {
	return b.v.status
}

func (b *Bike) GetChargesPerDay() int {
	return b.v.chargePerDay
}

func (b *Bike) SetStatus(status enum.VehicleStatus)  {
	b.v.status = status
}

func (b *Bike) GetOwnerName() string {
	return b.v.ownerName
}

func (b *Bike) GetVehicleType() enum.VechileType {
	return b.v.vechicleType
}