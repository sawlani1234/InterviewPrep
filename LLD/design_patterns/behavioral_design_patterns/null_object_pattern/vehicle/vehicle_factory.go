package vehicle


type VehicleFactory struct {

}


func NewVehicleFactory() VehicleFactory {
	return VehicleFactory{}
}


func (v VehicleFactory) GetVehicleFactory(vechicleType string ) Vehicle {

	switch vechicleType {
	case "CAR":
		return Car{}
	case "BIKE":
		return Bike{}
	}
	return  Null{}
}