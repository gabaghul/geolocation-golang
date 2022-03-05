package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gabaghul/geolocation-golang/consts"
	"github.com/gabaghul/geolocation-golang/models"
)

// GetCityInfoByName returns info from requested city
func GetCityInfoByName(cityName string) (cities models.GeoAPICities, err error) {
	client := consts.GetHTTPClient()
	headers := http.Header{}

	req, err := http.NewRequest("GET", consts.CitiesEndpoint, nil)
	if err != nil {
		return
	}

	queryStrings := req.URL.Query()

	queryStrings.Add("namePrefix", cityName)
	headers.Add("X-RapidAPI-Key", consts.GetApiKey())

	req.URL.RawQuery = queryStrings.Encode()
	req.Header = headers

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cities)

	return
}

// GetNearbyCities returns all nearby cities based on a city reference ID
func GetNearbyCities(cityID int64) (cities models.GeoAPICities, err error) {
	client := consts.GetHTTPClient()
	headers := http.Header{}

	req, err := http.NewRequest("GET", consts.GetNearbyCitiesEndpoint(cityID), nil)
	if err != nil {
		return
	}

	headers.Add("X-RapidAPI-Key", consts.GetApiKey())
	req.Header = headers

	res, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &cities)

	return
}
