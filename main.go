package main

import (
	"os"
	"whatsapp_service/routes"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	/*
		LOGRUS INIT
	*/
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	println("App started")
	err := godotenv.Load()
	if err != nil {
		return
	}
	routes.Init()
}
