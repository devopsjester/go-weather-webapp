package main

import (
	"log"

	"go-weather-webapp/internal/server"
	"go-weather-webapp/pkg/weather"
)

type PageData struct {
	Location *weather.Location
	Weather  *weather.Weather
	Error    string
}

func main() {
	srv := server.NewServer()
	srv.SetupRoutes()

	log.Println("Server is running on port 8080...")
	log.Fatal(srv.Start(":8080"))
}
