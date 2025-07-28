package theatre

import "bookmyshow/audi"

type Theatre struct {
	Id int
	Name string
	Audis []audi.Audi 
}

func NewTheatre(Id int,Name string,Audis []audi.Audi) *Theatre {
	return &Theatre{Id,Name,Audis}
}