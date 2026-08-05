package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvironmentVariables()

	r := router.Gerar()

	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Your application is running at port http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
