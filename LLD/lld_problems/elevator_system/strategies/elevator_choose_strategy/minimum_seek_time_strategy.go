package elevatorchoosestrategy

import "solid_design/elevator_system/enum"

type MinimumSeekTimeStrategy struct {

}

func (m MinimumSeekTimeStrategy) SubmitNewRequest(floor int,direction enum.Direction) int {
	return -1
}