package routes

import (
	"banking/apiServer/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterCustomerRoutes(router *mux.Router, handler *handlers.CustomersHandler) {
	registerCustomerGetRoutes(router, handler)
	registerCustomerPostRoutes(router, handler)
}

func registerCustomerGetRoutes(router *mux.Router, handler *handlers.CustomersHandler) {
	router.HandleFunc("/customers", handler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", handler.GetCustomerById).Methods(http.MethodGet)
}

func registerCustomerPostRoutes(router *mux.Router, handler *handlers.CustomersHandler) {

}
