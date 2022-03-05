package services

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gabaghul/geolocation-golang/api"
	"github.com/gabaghul/geolocation-golang/consts"
	"github.com/gabaghul/geolocation-golang/logger"
	"github.com/gabaghul/geolocation-golang/models"
)

func GetNearbyCities(ctx context.Context, cityName string) (res []byte, err error) {
	log := logger.GetLogger()

	cities, err := api.GetCityInfoByName(cityName)
	if err != nil {
		log.Err(err).Msgf(consts.DefaultErrorMessage)
		res = []byte(consts.DefaultErrorMessage)
		return
	}

	responseCities := make([]models.Cities, len(cities.Data))
	time.Sleep(time.Second * 1)
	for i, city := range cities.Data {
		nearby, errFor := api.GetNearbyCities(city.Id)
		if errFor != nil {
			log.Err(errFor).Msgf(consts.DefaultErrorMessage)
			res = []byte(consts.DefaultErrorMessage)
			return
		}

		var totalPopulation int64
		for _, n := range nearby.Data {
			totalPopulation += n.Population
		}

		responseCities[i] = models.Cities{
			ReferenceCity:    strings.Join([]string{city.City, city.Country}, "-"),
			NearbyPopulation: totalPopulation,
			NearbyCities:     nearby.Data,
		}
		time.Sleep(time.Second * 2)
	}

	return json.Marshal(&responseCities)
}
