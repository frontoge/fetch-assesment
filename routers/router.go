package routers

import (
	"fetch/receipt-processor/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", controllers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controllers.GetPoints).Methods("GET")

	return router
}
