package movie


type Movie struct {
	Id int 
	Name string 
	Genre string 
}


func NewMovie(id int,name,genre string) *Movie{
	return &Movie{id,name,genre}
}

