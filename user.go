package faker

import (
	"encoding/json"
	"fmt"
)

// API endpoint for random users.
const URL = API + "&inc=gender,name,email,dob,phone,cell,picture,nat"

// Returns a random user data.
func RandomUser() (*User, error) {
	response, err := JSONResponse(URL)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random users data.
func MultiRandomUser(count int) ([]User, error) {
	url := URL + fmt.Sprintf("&results=%d", count)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}

// Returns a random user by nationality.
func RandomUserByNationality(nationality []Nationality) (*User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&nat=%s", nat)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random users by nationality.
func MultiRandomUserByNationality(nationality []Nationality, count int) ([]User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&nat=%s&results=%d", nat, count)
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
	url := URL + "&gender=male"
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random male users data.
func MultiRandomMaleUser(count int) ([]User, error) {
	url := URL + fmt.Sprintf("&gender=male&results=%d", count)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}

// Returns a random male user by nationality.
func RandomMaleUserByNationality(nationality []Nationality) (*User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&gender=male&nat=%s", nat)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random male users by nationality.
func MultiRandomMaleUserByNationality(nationality []Nationality, count int) ([]User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&gender=male&nat=%s&results=%d", nat, count)
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
	url := URL + "&gender=female"
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random female users data.
func MultiRandomFemaleUser(count int) ([]User, error) {
	url := URL + fmt.Sprintf("&gender=female&results=%d", count)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}

// Returns a random male user by nationality.
func RandomFemaleUserByNationality(nationality []Nationality) (*User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&gender=female&nat=%s", nat)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return &(data["results"][0]), nil
}

// Returns multiple random male users by nationality.
func MultiRandomFemaleUserByNationality(nationality []Nationality, count int) ([]User, error) {
	nat := NatString(nationality)
	url := URL + fmt.Sprintf("&gender=female&nat=%s&results=%d", nat, count)
	response, err := JSONResponse(url)
	if err != nil {
		return nil, err
	}
	var data map[string][]User
	json.Unmarshal(response, &data)
	return data["results"], nil
}
