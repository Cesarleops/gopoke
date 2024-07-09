package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

//Build a struct to work with the response of the poke api.
//The json:name syntax function is to map the json response to a struct

type RespLocations struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) ListPokemons(pageUrl *string) (RespLocations, error) {
	url := baseURL + "location-area"
	cache := NewCache(5 * time.Second)

	if pageUrl != nil {
		url = *pageUrl
	}

	if v, ok := cache.Get(url); ok {
		locations := RespLocations{}
		err := json.Unmarshal(v, &locations)
		if err != nil {
			return RespLocations{}, nil
		}
		return locations, nil
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		//The zero value of a struct is the empty struct.
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespLocations{}, nil
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespLocations{}, err
	}

	myLocations := RespLocations{}
	err = json.Unmarshal(data, &myLocations)
	cache.Add(url, data)
	if err != nil {
		return RespLocations{}, err
	}

	return myLocations, nil
}
