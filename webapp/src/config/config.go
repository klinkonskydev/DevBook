package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the URL used to communicate with the API
	APIURL = ""
	// Port is the port where the web application is running
	Port = 0
	// HashKey is used to authenticate the cookie
	HashKey []byte
	// BlockKey is used to encrypt the cookie data
	BlockKey []byte
)

// LoadEnvironmentVariables initializes the environment variables
func LoadEnvironmentVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
