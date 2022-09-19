package main

import (
	"log"
	"steamsale/api"
	"steamsale/bootstrap"
	"steamsale/repository"
	"steamsale/service"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("Problem with config load: ", err)
	}

	errorList := cfg.Validate()
	if errorList != nil {
		log.Fatalln("Problem with config validation: ", errorList)
	}

	db, err := bootstrap.DBConnect(cfg)
	if err != nil {
		log.Fatalln("Problem with Database connection: ", err)
	}

	repo := repository.New(db)

	userService, itemService := service.New(repo)

	server := api.NewServer(cfg.HTTPPort)

	err = server.Start(userService, itemService)
	if err != nil {
		log.Fatalln("Problem with server startup: ", err)
	}
}
