package faker

import (
	"io/ioutil"
	"net/http"
)

const API = "https://randomuser.me/api?noinfo"

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
