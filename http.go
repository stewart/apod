package apod

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint = "https://api.nasa.gov/planetary/apod"

func get(params map[string]string) (*APOD, error) {
	var apod = &APOD{}

	route, _ := url.Parse(endpoint)
	query := url.Values{}

	for key, value := range params {
		query.Add(key, value)
	}

	route.RawQuery = query.Encode()

	res, err := http.Get(route.String())
	if err != nil {
		return apod, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return apod, err
	}

	if err := json.Unmarshal(data, apod); err != nil {
		return apod, err
	}

	return apod, nil
}
