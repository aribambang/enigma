package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name       string `json:"name"`
	City       string `json:"city"`
	PostalCode int    `json:"postal_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ari", "Bekasi", 17411},
		{"Bambang", "Bekasi", 17156},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
