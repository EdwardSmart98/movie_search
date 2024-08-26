package main

import (
	"github.com/joho/godotenv"
	"log"
	"movieInfo/cmd/apiServer"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	apiServer.Run()

}
