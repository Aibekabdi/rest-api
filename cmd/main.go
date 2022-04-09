package main

import (
	"log"

	restapi "github.com/Aibekabdi/rest-api"
	"github.com/Aibekabdi/rest-api/internal/delivery"
	"github.com/Aibekabdi/rest-api/internal/repository"
	"github.com/Aibekabdi/rest-api/internal/service"
)

func main() {

	//preparation
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := delivery.NewHandler(services)
	//running server
	log.Println("Server is listenning...")
	log.Println("http://localhost:8080/")
	srv := new(restapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
