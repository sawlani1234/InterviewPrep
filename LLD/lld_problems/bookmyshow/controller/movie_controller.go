package controller

import (
	"bookmyshow/enum"
	"bookmyshow/movie"
)


type MovieController struct {
	MovieCityMap map[enum.CityType][]movie.Movie
}

func NewMovieController() *MovieController{

  return &MovieController{MovieCityMap: make(map[enum.CityType][]movie.Movie,0)}	

}

func (m *MovieController) AddMovie(city enum.CityType,movie movie.Movie) {
	m.MovieCityMap[city] = append(m.MovieCityMap[city], movie)
}