package main

import (
	"golang-jwt/config"
	"golang-jwt/lib/env"
	"log"
)


func main() {
	err := config.Routers.Run(env.String("MainSetup.ServerHost", ""))
	if err != nil {
		log.Fatal(err)
		return 
	}
}