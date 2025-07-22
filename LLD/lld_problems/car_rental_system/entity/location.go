package entity


type Location struct {
	id      int
	pincode int
	address string
}

func NewLocation(pincode int, address string) Location{
	return Location{
		id: 1,
		pincode: pincode,
		address: address,
	}
}

func (l Location) GetPincode() int {
	return l.pincode
}