package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
  // DataSourceName is a connection string for mysql-server
	DataSourceName = ""
  // Port is a port where the API will be running
	Port       = 0
)

// LoadEnvironmentVariables will load the environment variables
func LoadEnvironmentVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		// Fatal because this error is an API error and not a Request error.
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 5000
	}

	DataSourceName = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
