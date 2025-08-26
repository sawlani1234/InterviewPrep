package auction 


type IplAuction struct {
	collegue []Collegue
}


func NewIplAuction() *IplAuction {
	return &IplAuction{}
}


func (i *IplAuction) addBidder(c Collegue) {
	i.collegue = append(i.collegue, c)
}


func (i *IplAuction) addBid(c Collegue,amount int) {
	
	for _,val := range i.collegue {
		if val.getId() != c.getId() {
			val.recieveNotifcation(amount)
		}
	}
}
