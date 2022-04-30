package faker

import (
	"encoding/json"
	"fmt"
)

// Returns a random picture.
func RandomPicture() (*Picture, error) {
	response, err := JSONResponse(API)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	json.Unmarshal(response, &data)
	return &(data["results"][0].Picture), nil
}

// Returns multiple random pictures.
func MultiRandomPicture(results int) ([]Picture, error) {
	url := API + fmt.Sprintf("&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	var result []Picture
	json.Unmarshal(response, &data)
	for _, value := range data["results"] {
		result = append(result, value.Picture)
	}
	return result, nil
}
