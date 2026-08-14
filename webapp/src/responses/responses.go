package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError represents an error response coming from the API
type APIError struct {
	Error string `json:"error"`
}

// JSON returns a JSON response for the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// HandleStatusCodeError handles requests with status code 400 or above
func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var apiError APIError
	json.NewDecoder(r.Body).Decode(&apiError)
	JSON(w, r.StatusCode, apiError)
}
