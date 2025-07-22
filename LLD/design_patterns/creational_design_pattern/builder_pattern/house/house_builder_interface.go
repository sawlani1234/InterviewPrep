package house


type HouseBuilder interface {
	SetHouseName(name string) HouseBuilder
	SetDoors(n int) HouseBuilder
	SetWalls(n int) HouseBuilder
	Build() House
}