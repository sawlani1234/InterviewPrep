package auction

type Collegue interface {
	AddBid(amount int)
	getId() int
	recieveNotifcation(amount int)
}
