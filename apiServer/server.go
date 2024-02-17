package apiServer

import (
	"banking/apiServer/handlers"
	"banking/apiServer/routes"
	"banking/logger"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiServer struct {
	host            string
	port            string
	customerHandler *handlers.CustomersHandler
	accountHandler  *handlers.AccountsHandler
}

func NewApiServer(
	host string,
	port string,
	customerHandler *handlers.CustomersHandler,
	accountHandler *handlers.AccountsHandler) *ApiServer {
	return &ApiServer{
		host,
		port,
		customerHandler,
		accountHandler,
	}
}

func (server *ApiServer) Start() error {
	//mux := http.NewServeMux()
	logger.Info(fmt.Sprintf("Starting Server at host %v and port %v", server.host, server.port))
	router := mux.NewRouter()

	server.registerRoutes(router)

	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", server.host, server.port), router); err != nil {
		return err
	}
	return nil
}

func (server *ApiServer) registerRoutes(router *mux.Router) {
	//customer routes
	routes.RegisterCustomerRoutes(router, server.customerHandler)
	//account routes
	routes.RegisterAccountRoutes(router, server.accountHandler)
}
