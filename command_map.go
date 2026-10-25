package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(conf *config) error {
	//res, err := http.Get("https://pokeapi.co/api/v2/location-area/{id or name}/")
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(locationsResp)
	return err
}

