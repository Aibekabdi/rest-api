package main

import (
	"context"
	"log"

	restapi "github.com/Aibekabdi/rest-api"
	"github.com/Aibekabdi/rest-api/internal/delivery"
	"github.com/Aibekabdi/rest-api/internal/repository"
	"github.com/Aibekabdi/rest-api/internal/service"
	"github.com/go-redis/redis/v8"
)

func main() {
	//redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	log.Fatalf("the redis db has error while connecting: %s", err.Error())
	//preparation
	repos := repository.NewRepository(client)
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
