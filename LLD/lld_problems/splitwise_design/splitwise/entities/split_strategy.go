package entities


type SplitStrategy interface {
     Split(amount int,splits []Split) ([]Split,error)
}




