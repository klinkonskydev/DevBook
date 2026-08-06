package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response for any request
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)

  if err := json.NewEncoder(w).Encode(data); err != nil {
    log.Fatal(err)
  }
}

// Error return an error in JSON
func Error(w http.ResponseWriter, statusCode int, error error) {
  JSON(w, statusCode, struct {
    Error string `json:"error"`
  }{
    Error: error.Error(),
  })
}
