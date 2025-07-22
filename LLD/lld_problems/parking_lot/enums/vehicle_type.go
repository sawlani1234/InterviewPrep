package enums

type VehicleType int

const (
	Wheeler VehicleType = iota
	TwoWheeler
	FourWheeler
)

var MapIntToVehicleType = map[int]VehicleType{
	1: TwoWheeler,
	2: FourWheeler,
}
