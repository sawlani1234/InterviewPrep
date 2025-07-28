package controller

import (
	"bookmyshow/show"
)


type ShowController struct {
	AudiToShowMap map[int][]show.Show
}

func NewShowController() *ShowController {
	return &ShowController{AudiToShowMap: make(map[int][]show.Show,0)}
}


func (s *ShowController) AddShows(audiId int,show show.Show) {
	s.AudiToShowMap[audiId] = append(s.AudiToShowMap[audiId], show)
}