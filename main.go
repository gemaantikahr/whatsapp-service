package main

import (
	"whatsapp_service/routes"

	"github.com/joho/godotenv"
)

func main() {
	println("App started")
	godotenv.Load()
	routes.Init()
}
