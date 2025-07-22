package entity

type Store struct {
     id int
	 reservationManager *ReservationManager
	 vehicleManager *VehicleManager
	 location Location

}

func NewStore(id int,reservationManager *ReservationManager,vehicleManager *VehicleManager,location Location) Store {
	return Store{
		id: id,
		reservationManager: reservationManager,
		vehicleManager: vehicleManager,
		location: location,
	}
}

func (s Store) GetId() int {
	return s.id
}

func (s Store) GetLocation() Location {
	return s.location
}


func (s Store) GetVehicleManager() *VehicleManager {
	return s.vehicleManager
}

func (s Store) GetReservationManager() *ReservationManager {
	return s.reservationManager
}