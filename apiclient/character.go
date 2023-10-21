package apiclient

import "time"

type Character struct {
	Name         string `json:"name"`
	Height       string `json:"height"`
	Mass         string `json:"mass"`
	HairColor    string `json:"hair_color"`
	SkinColor    string `json:"skin_color"`
	EyeColor     string `json:"eye_color"`
	BirthYear    string `json:"birth_year"`
	Gender       string `json:"gender"`
	Homeworld    string `json:"homeworld"`
	HomePlanet   *HomePlanet
	FilmUrls     []string `json:"films"`
	SpeciesUrls  []string `json:"species"`
	Species      *Species
	VehicleUrls  []string `json:"vehicles"`
	StarshipUrls []string `json:"starships"`
	Starships    []Starship
	Created      time.Time `json:"created"`
	Edited       time.Time `json:"edited"`
	URL          string    `json:"url"`
}
