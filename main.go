package main

import (
	"context"

	"github.com/gabaghul/geolocation-golang/logger"
	"github.com/gabaghul/geolocation-golang/server"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Panic().Msg("PANIC! couldnt load environment variables, do you have a .env file created?")
	}

	log.Info().Msg("RUNNING, LET'S GO!")
	log.Info().Msg("To kill this process, just CTRL+C and i will be interrupted")

	ctx := context.Background()
	server.Start(ctx)
}
