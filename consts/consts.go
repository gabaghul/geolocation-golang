package consts

import (
	"os"
	"sync"
)

var (
	once   sync.Once
	apiKey string
)

func GetApiKey() string {
	once.Do(func() {
		apiKey = os.Getenv("GEOAPI_API_KEY")
	})

	return apiKey
}
