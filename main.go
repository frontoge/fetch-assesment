package main

import (
	"log"
	"net/http"
	"fetch/receipt-processor/routers"
)

func main() {
	router := routers.InitRouter()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}