package faker

import (
	"encoding/json"
	"fmt"
)

// Returns a random location.
func RandomLocation() (*Location, error) {
	response, err := JSONResponse(API)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	json.Unmarshal(response, &data)
	return &(data["results"][0].Location), nil
}

// Returns multiple random locations.
func MultiRandomLocation(results int) ([]Location, error) {
	url := API + fmt.Sprintf("&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	var result []Location
	json.Unmarshal(response, &data)
	for _, value := range data["results"] {
		result = append(result, value.Location)
	}
	return result, nil
}
