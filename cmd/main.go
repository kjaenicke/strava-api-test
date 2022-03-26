package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kjaenicke/strava-api-test/internal/server"
)

func main(){
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := server.CreateServer()
	port := os.Getenv("PORT")

	fmt.Println("Starting strava-api-test on port " + port)
	server.Run(":" + port)
}