package enum


type CityType string


const (

	Bengaluru CityType = "Bengaluru"

	Delhi CityType = "Delhi"
)


func GetAllCities() []CityType{
   return []CityType{Bengaluru,Delhi}
}