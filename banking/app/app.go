package app

import (
	"log"
	"net/http"

	"github.com/aribambang/enigma/banking/domain"
	"github.com/aribambang/enigma/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
