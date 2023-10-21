package report

import (
	"fmt"
	api "swapi_go/apiclient"
	"time"
)

type Report interface {
	PrintCharacters(characters []api.Character) error
}

type StarWarsReport struct {
	// TODO - Add output options for report
}

func (r StarWarsReport) PrintCharacters(characters []api.Character) error {
	fmt.Println("Search Results: ", time.Now())
	fmt.Println("Characters found: ", len(characters))
	for chararacterIndex, character := range characters {
		fmt.Println("------------------------------------------------")
		fmt.Println("Character #", chararacterIndex+1)
		fmt.Println("Name: ", character.Name)
		if len(character.Starships) > 0 {
			for index, starship := range character.Starships {
				fmt.Println("Starship", index+1)
				fmt.Println("    Name: ", starship.Name)
				fmt.Println("    Cargo Capacity: ", starship.CargoCapacity)
				fmt.Println("    Class: ", starship.StarshipClass)
			}
		} else {
			fmt.Println("Starships: NA")
		}

		if character.Species != nil {
			fmt.Println("Species")
			fmt.Println("    Name: ", character.Species.Name)
			fmt.Println("    Language: ", character.Species.Language)
			fmt.Println("    Lifespan: ", character.Species.AverageLifespan)
		} else {
			fmt.Println("Species NA")
		}

		if character.HomePlanet != nil {
			fmt.Println("Home Planet")
			fmt.Println("    Name: ", character.HomePlanet.Name)
			fmt.Println("    Population: ", character.HomePlanet.Population)
			fmt.Println("    Climate: ", character.HomePlanet.Climate)
		} else {
			fmt.Println("Home Planet NA")
		}
		fmt.Println("------------------------------------------------")
	}

	return nil
}
