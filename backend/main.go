package main

import (
	"log"
	"telnesstech/api"
	router "telnesstech/router"
)

func main() {
	service, err := api.NewService()
	if err != nil {
		log.Fatalf("error while running service")
	}
	router.Run("localhost:8080", service.Handler)
}
