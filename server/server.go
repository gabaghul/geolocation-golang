package server

import (
	"context"
	"net/http"
	"os"

	"github.com/gabaghul/geolocation-golang/logger"
	"github.com/gabaghul/geolocation-golang/services"
	"github.com/gorilla/mux"
)

// Start starts our HTTP server
func Start(ctx context.Context) {
	log := logger.GetLogger()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("CALLED /")
		w.Write([]byte("How you're doing? cuz im running just fine! (:"))
	}).Methods("GET")

	r.HandleFunc("/nearbyCities/{cityName}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		cityName := params["cityName"]
		log.Debug().Msgf("CALLED /nearbyCities/%s", params["cityName"])

		res, err := services.GetNearbyCities(ctx, cityName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	port := os.Getenv("PORT")
	log.Debug().Msgf("Running on port %s", port)

	http.ListenAndServe(":"+port, r)
}
