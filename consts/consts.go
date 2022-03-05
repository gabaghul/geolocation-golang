package consts

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	rootEndpoint   = "https://wft-geo-db.p.rapidapi.com/v1/geo"
	CitiesEndpoint = rootEndpoint + "/cities"
)

var (
	once   sync.Once
	apiKey string
	client *http.Client
)

func GetApiKey() string {
	once.Do(func() {
		apiKey = os.Getenv("GEOAPI_API_KEY")
	})

	return apiKey
}

func GetHTTPClient() *http.Client {
	once.Do(func() {
		client = &http.Client{}
	})

	return client
}

func GetNearbyCitiesEndpoint(id int64) string {
	return fmt.Sprintf(CitiesEndpoint+"/%d/nearbyCities", id)
}
