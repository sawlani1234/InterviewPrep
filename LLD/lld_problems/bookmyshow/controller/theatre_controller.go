package controller

import (
	"bookmyshow/enum"
	"bookmyshow/theatre"
)


type TheatreController struct {
	CityToTheatreMap map[enum.CityType][]theatre.Theatre
}

func NewTheatreController() *TheatreController {
	return &TheatreController{CityToTheatreMap: make(map[enum.CityType][]theatre.Theatre,0)}
}

func (t *TheatreController) AddTheatre(cityType enum.CityType,theatre theatre.Theatre) {
	t.CityToTheatreMap[cityType] = append(t.CityToTheatreMap[cityType],theatre)
}