package controllers

import (
	"encoding/json"
	"fetch/receipt-processor/helpers"
	"fetch/receipt-processor/models"
	"fetch/receipt-processor/scoring"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// parse request body
	var parsedRequest models.ProcessReceiptRequest

	if err := json.NewDecoder(r.Body).Decode(&parsedRequest); err != nil {
		log.Default().Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate request object
	validationErr := helpers.ValidateProcessRequest(parsedRequest)

	if validationErr != nil {
		log.Default().Printf("Error validating request: %v", validationErr)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// map request object types
	data, mapErr := helpers.MapProcessRequest(parsedRequest)

	if mapErr != nil {
		log.Default().Printf("Error mapping request: %v", mapErr)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Run scoring algorithm and generate UUID for receipt
	id := scoring.ProcessReceipt(data)

	// send the response
	response := models.ProcessReceiptResponse{Id: id}
	json.NewEncoder(w).Encode(response)
	return
}

// Gets points for a receipt given a UUID
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	score := scoring.GetScoreById(vars["id"])

	// scoring getter returns -1 if receipt not found in memory with given ID
	if score == -1 {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	} else {
		response := models.PointsResponse{Points: score}
		json.NewEncoder(w).Encode(response)
		return
	}
}
