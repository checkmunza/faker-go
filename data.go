package faker

import (
	"encoding/json"
	"fmt"
)

// Returns random user data, which includes user information,
// location, login details and pictures.
func RandomData() (*Data, error) {
	response, err := JSONResponse(API)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	json.Unmarshal(response, &data)
	return &data["results"][0], nil
}

// Returns multiple random users data.
func MultiRandomData(count int) ([]Data, error) {
	url := API + fmt.Sprintf("&results=%d", count)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	json.Unmarshal(response, &data)
	return data["results"], nil
}
