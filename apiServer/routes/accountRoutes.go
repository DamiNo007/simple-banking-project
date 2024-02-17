package routes

import (
	"banking/apiServer/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAccountRoutes(router *mux.Router, handler *handlers.AccountsHandler) {
	registerAccountGetRoutes(router, handler)
	registerAccountPostRoutes(router, handler)
}

func registerAccountGetRoutes(router *mux.Router, handler *handlers.AccountsHandler) {

}

func registerAccountPostRoutes(router *mux.Router, handler *handlers.AccountsHandler) {
	router.HandleFunc("/customers/{customerId:[0-9]+}/account", handler.NewAccount).Methods(http.MethodPost)
}
