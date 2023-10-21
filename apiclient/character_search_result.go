package apiclient

type CharacterSearchResult struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Character `json:"results"`
}
