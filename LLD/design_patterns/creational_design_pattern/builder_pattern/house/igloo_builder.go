package house


type IglooBuilder struct {
	name string 
	door int 
	walls int
}

func NewIglooBuilder() IglooBuilder{
	return IglooBuilder{}
}

func (i IglooBuilder) SetHouseName(name string) HouseBuilder {
	i.name = name 
	return i
}

func (i IglooBuilder) SetDoors(n int) HouseBuilder {
     i.door = n
	 return i 
}

func (i IglooBuilder) SetWalls(n int) HouseBuilder {
	i.walls = n 
	return  i 
}

func (i IglooBuilder) Build() House {
	return House{ 
		i.name,
		i.door,
		i.walls,
	}
}


