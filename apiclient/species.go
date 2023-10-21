package apiclient

import "time"

type Species struct {
	Name            string    `json:"name"`
	Classification  string    `json:"classification"`
	Designation     string    `json:"designation"`
	AverageHeight   string    `json:"average_height"`
	SkinColors      string    `json:"skin_colors"`
	HairColors      string    `json:"hair_colors"`
	EyeColors       string    `json:"eye_colors"`
	AverageLifespan string    `json:"average_lifespan"`
	Homeworld       string    `json:"homeworld"`
	Language        string    `json:"language"`
	People          []string  `json:"people"`
	Films           []string  `json:"films"`
	Created         time.Time `json:"created"`
	Edited          time.Time `json:"edited"`
	URL             string    `json:"url"`
}
