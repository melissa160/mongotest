package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liftitapp/mongotest/handlers"
	"github.com/liftitapp/mongotest/utils"
)

func main() {

	// Enviroment variables validation
	err := utils.ValidateEnvVars()
	if err != nil {
		log.Fatal(err)
	}
	// Charge error code array
	utils.SetupErrorCode()

	// DB connection
	db, err := utils.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/tracking", handlers.CreateTrackerRegister).Methods("POST")
	//router.HandleFunc("/tracking/{id}", handlers.GetTrackerRegister).Methods("GET")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
