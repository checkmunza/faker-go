package faker

import (
	"encoding/json"
	"fmt"
)

// Returns a random user data.
func RandomUser() (*User, error) {
	url := API + "&inc=gender,name,email,dob,phone,cell"
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random users data.
func MultiRandomUser(results int) ([]User, error) {
	url := API + fmt.Sprintf("&inc=gender,name,email,dob,phone,cell&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}

// Returns a random male user's data.
func RandomMaleUser() (*User, error) {
	url := API + "&inc=gender,name,email,dob,phone,cell&gender=male"
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random male users data.
func MultiRandomMaleUser(results int) ([]User, error) {
	url := API + fmt.Sprintf("&inc=gender,name,email,dob,phone,cell&gender=male&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}

// Returns a random female user's data.
func RandomFemaleUser() (*User, error) {
	url := API + "&inc=gender,name,email,dob,phone,cell&gender=female"
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random female users data.
func MultiRandomFemaleUser(results int) ([]User, error) {
	url := API + fmt.Sprintf("&inc=gender,name,email,dob,phone,cell&gender=female&results=%d", results)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}
