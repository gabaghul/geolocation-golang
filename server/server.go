package server

import (
	"context"
	"net/http"
	"os"

	"github.com/gabaghul/geolocation-golang/logger"
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

	port := os.Getenv("PORT")
	log.Debug().Msgf("Running on port %s", port)

	http.ListenAndServe(":"+port, r)
}
