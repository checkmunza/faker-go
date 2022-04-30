package faker

import (
	"io/ioutil"
	"net/http"
)

const API = "https://randomuser.me/api?noinfo"

type Nationality int

// Enum for nationalities.
const (
	AU Nationality = iota
	BR
	CA
	CH
	DE
	DK
	ES
	FI
	FR
	GB
	IE
	IR
	NO
	NL
	NZ
	TR
	US
)

// Returns value of enum type.
func (n Nationality) Value() string {
	return [...]string{"AU", "BR", "CA", "CH", "DE", "DK", "ES", "FI", "FR", "GB", "IE", "IR", "NO", "NL", "NZ", "TR", "US"}[n]
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
