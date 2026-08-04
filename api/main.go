package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

var port string = ":5000"

func main() {
	r := router.Gerar()

	fmt.Printf("Your application is running at port http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
