package ball


type Ball struct {
	runs int 
	isExtra bool 

}


func NewBall(runs int,isExtra bool) *Ball {
	return &Ball{runs,isExtra}
}