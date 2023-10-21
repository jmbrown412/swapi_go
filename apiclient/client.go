package apiclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type StarWarsClient interface {
	SearchCharacters(characterName string) ([]Character, error)
	GetStarship(url string) (*Starship, error)
	GetHomePlanet(url string) (*HomePlanet, error)
	GetSpecies(url string) (*Species, error)
}

type SwapiApiClient struct {
	url string
}

func NewSwapiApiClient(url string) SwapiApiClient {
	return SwapiApiClient{url: url}
}

func (c SwapiApiClient) SearchCharacters(characterName string) ([]Character, error) {
	var characters []Character
	searchResult, err := c.searchCharacters(characterName, nil)
	if err != nil {
		return nil, err
	}
	characters = append(characters, searchResult.Results...)

	for searchResult.Next != "" {
		searchResult, err = c.searchCharacters(characterName, &searchResult.Next)
		if err != nil {
			return nil, err
		}
		characters = append(characters, searchResult.Results...)
	}

	// Sort characters alphabetically ascending
	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Name < characters[j].Name
	})
	return characters, nil
}

func (c SwapiApiClient) GetStarship(url string) (*Starship, error) {
	var starship *Starship
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET starship request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("GET starship request failed with status code: ", resp.StatusCode)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&starship)
	if err != nil {
		fmt.Println("Error decoding JSON from GET starship request:", err)
		return nil, err
	}
	return starship, nil
}

func (c SwapiApiClient) GetHomePlanet(url string) (*HomePlanet, error) {
	var homePlanet *HomePlanet
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET planet request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("GET planet request failed with status code: ", resp.StatusCode)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&homePlanet)
	if err != nil {
		fmt.Println("Error decoding JSON from GET planet request:", err)
		return nil, err
	}
	return homePlanet, nil
}

func (c SwapiApiClient) GetSpecies(url string) (*Species, error) {
	var species *Species
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the GET species request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("GET species request failed with status code: ", resp.StatusCode)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&species)
	if err != nil {
		fmt.Println("Error decoding JSON from GET species request:", err)
		return nil, err
	}
	return species, nil
}

func (c SwapiApiClient) searchCharacters(characterName string, nextUrl *string) (*CharacterSearchResult, error) {
	var result *CharacterSearchResult

	// Setup the searchurl. It can be the next page url
	var searchUrl string
	if nextUrl == nil {
		searchUrl = fmt.Sprintf("%s/people?search=%s", c.url, characterName)
	} else {
		searchUrl = *nextUrl
	}

	resp, err := http.Get(searchUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code: ", resp.StatusCode)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// Hydrate Character details from other resources since API
	// returns urls for things like Planet, Starships, Species, etc...
	for charIndex, character := range result.Results {
		// Hydrate Species
		if len(character.SpeciesUrls) > 0 {
			species, err := c.GetSpecies(character.SpeciesUrls[0])
			if err != nil {
				return nil, err
			}
			character.Species = species
		}

		// Hydrate Starships
		if len(character.StarshipUrls) > 0 {
			for _, starshipUrl := range character.StarshipUrls {
				starship, err := c.GetStarship(starshipUrl)
				if err != nil {
					return nil, err
				}
				character.Starships = append(character.Starships, *starship)
			}
		}

		// Hydrate Planets
		if len(character.Homeworld) > 0 {
			if len(character.SpeciesUrls) > 0 {
				homePlanet, err := c.GetHomePlanet(character.Homeworld)
				if err != nil {
					return nil, err
				}
				character.HomePlanet = homePlanet
			}
		}

		// Finally, update the Character
		result.Results[charIndex] = character
	}

	return result, nil
}
