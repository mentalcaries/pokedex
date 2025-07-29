package pokeapi

type LocationData struct {
    Name string `json:"name"`
    URL string  `json:"url"`
}

type LocationResult struct {
    Count int `json:"count"`
    Next *string  `json:"next"`
    Previous *string `json:"previous"`
    Results []LocationData `json:"results"`
}