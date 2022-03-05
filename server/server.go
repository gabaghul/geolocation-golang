package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start starts our HTTP server
func Start(ctx context.Context) {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("CALLED /")
		w.Write([]byte("How you're doing? cuz im running just fine! (:"))
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
