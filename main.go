package main

import (
	"Backend-Go/config"
	"Backend-Go/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    config.ConnectDB()

    r := routes.SetupRouter()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
