package main

import (
	"fmt"
	api "swapi_go/apiclient"
	report "swapi_go/report"
)

func main() {
	fmt.Println("Star Wars Character Search v1.0")

	// Setup the necessary services
	apiClient := api.NewSwapiApiClient("https://swapi.dev/api")
	report := report.StarWarsReport{}

	// Get the characterName from the user
	var characterName string
	fmt.Println("Enter the Star Wars character name: ")
	_, err := fmt.Scan(&characterName)
	if err != nil {
		fmt.Println("There was an error getting the character input: ", err)
		return
	}
	fmt.Println("You entered:", characterName)

	// Search for the character(s)
	characters, err := apiClient.SearchCharacters(characterName)
	if err != nil {
		fmt.Println("There was an error calling the API: ", err.Error())
		return
	}

	// Print the characters report
	err = report.PrintCharacters(characters)
}
