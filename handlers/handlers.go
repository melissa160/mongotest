package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/liftitapp/mongotest/utils"
)

func CreateTrackerRegister(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//payload := utils.CreateTracker()
	// for _, p := range payload {
	// 	if p.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(p)
	// 		return
	// 	}
	// }
	fmt.Printf("\n11111111111 %#v", params)
	json.NewEncoder(w).Encode("Tracker not found")
}
