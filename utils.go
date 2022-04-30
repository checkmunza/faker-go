package faker

import (
	"io/ioutil"
	"net/http"
)

const API = "https://randomuser.me/api?noinfo"

type Nationality string

// Enum for nationalities.
const (
	AU Nationality = "AU"
	BR Nationality = "BR"
	CA Nationality = "CA"
	CH Nationality = "CH"
	DE Nationality = "DE"
	DK Nationality = "DK"
	ES Nationality = "ES"
	FI Nationality = "FI"
	FR Nationality = "FR"
	GB Nationality = "GB"
	IE Nationality = "IE"
	IR Nationality = "IR"
	NO Nationality = "NO"
	NL Nationality = "NL"
	NZ Nationality = "NZ"
	TR Nationality = "TR"
	US Nationality = "US"
)

// Returns value of enum type.
func (n Nationality) Value() string {
	return string(n)
}

// Returns the JSON response fetched from an API.
func JSONResponse(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NatString(nationality []Nationality) string {
	var nat string
	for index, data := range nationality {
		nat += data.Value()
		if index != len(nationality)-1 {
			nat += ","
		}
	}
	return nat
}
