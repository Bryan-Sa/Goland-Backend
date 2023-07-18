package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, response interface{}, code int) {

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
