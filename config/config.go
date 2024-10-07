package config

import (
	"log"
	"os"
)

func GetMongoURI() string {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI environment variable is not set")
    }
    return uri
}
