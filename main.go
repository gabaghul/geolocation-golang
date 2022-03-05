package main

import (
	"context"
	"log"

	"github.com/gabaghul/geolocation-golang/server"
)

func main() {
	log.Println("RUNNING, LET'S GO!")
	log.Println("To kill this process, just CTRL+C and i will be interupted")

	ctx := context.Background()
	server.Start(ctx)
}
