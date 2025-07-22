package house 

type VillaBuilder struct {
	name string 
	doors int 
	walls int 
    
}

func NewVillaBuilder() VillaBuilder {
	return VillaBuilder{}
}

func (v VillaBuilder) SetHouseName(name string) HouseBuilder {
	v.name = name 
	return v
}

func (v VillaBuilder) SetDoors(n int) HouseBuilder {
	v.doors = n 
	return v
}

func (v VillaBuilder) SetWalls(n int) HouseBuilder {
	v.walls = n
	return  v
}

func (v VillaBuilder) Build() House {
	return House{
		v.name,
		v.doors,
		v.walls,
	}
}
