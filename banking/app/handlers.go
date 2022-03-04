package app

import (
	"encoding/json"
	"net/http"

	"github.com/aribambang/enigma/banking/service"
)

type Customer struct {
	Name       string `json:"name"`
	City       string `json:"city"`
	PostalCode int    `json:"postal_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
