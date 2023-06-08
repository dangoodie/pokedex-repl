package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResp{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
	return locationAreasResp, nil
}
