package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMapf(cfg *config, args ...string) error {
    url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocationsURL != nil {
    	url = *cfg.nextLocationsURL
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err}
	if res.StatusCode > 299 {
		return fmt.Errorf("got non-ok status code: %d", res.StatusCode)
	}
	locations := RespShallowLocations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	for _, location := range(locations.Results) {
		fmt.Println(location.Name)
	}
    return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	res, err := http.Get(*cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err}
	if res.StatusCode > 299 {
		return fmt.Errorf("got non-ok status code: %d", res.StatusCode)
	}
	locations := RespShallowLocations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	for _, location := range(locations.Results) {
		fmt.Println(location.Name)
	}
	return nil //
}

type RespShallowLocations struct {
    Count    int             `json:"count"`
    Next     *string         `json:"next"`
    Previous *string         `json:"previous"`
    Results  []ShallowLocation `json:"results"`
}

type ShallowLocation struct {
    Name 	string 	`json:"name"`
	URL 	string 	`json:"url"` 
}