package pokeapi

import (
	"encoding/json"
	"io"
	"errors"
	"net/http"
)



func (c *Client) GetPokemon(query string) (Pokemon, error) {
	endpoint := "pokemon/"
	fullURL := baseURL + endpoint + query

	// check cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := errors.New(resp.Status)
		return Pokemon{}, err
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// cache data
	c.cache.Add(fullURL, dat)

	return pokemon, nil
}
