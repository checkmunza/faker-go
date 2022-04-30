package faker

import (
	"encoding/json"
	"fmt"
)

// Returns a random login credential.
func RandomLogin() (*Login, error) {
	response, err := JSONResponse(API)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	json.Unmarshal(response, &data)
	return &(data["results"][0].Login), nil
}

// Returns multiple random login credentials.
func MultiRandomLogin(results int) ([]Login, error) {
	url := API + fmt.Sprintf("&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]Data
	var result []Login
	json.Unmarshal(response, &data)
	for _, value := range data["results"] {
		result = append(result, value.Login)
	}
	return result, nil
}
