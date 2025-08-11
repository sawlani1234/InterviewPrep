package weight


type KgAdapter struct {
	lbs LbsWeight
}


func NewKgAdapter(lbs LbsWeight) *KgAdapter {
	return &KgAdapter{lbs}
}

func (k *KgAdapter) GetWeightInKg() int {

	return k.lbs.GetWeightInLbs()*2
}