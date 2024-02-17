package handlers

import (
	"banking/apiServer/wrappers"
	"banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomersHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) CustomersHandler {
	return CustomersHandler{service: service}
}

func (h *CustomersHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	format := retrieveFormatFromRequestHeaders(r)

	if status := r.URL.Query().Get("status"); status != "" {
		customers, appErr := h.service.GetCustomersByStatus(status)
		handleApiResponse(w, format, wrappers.NewResponseWrapper(http.StatusOK, customers), appErr)
	} else {
		customers, appErr := h.service.GetAllCustomers()
		handleApiResponse(w, format, wrappers.NewResponseWrapper(http.StatusOK, customers), appErr)
	}
}

func (h *CustomersHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customerId"]

	customer, appErr := h.service.GetCustomerById(id)
	format := retrieveFormatFromRequestHeaders(r)

	handleApiResponse(w, format, wrappers.NewResponseWrapper(http.StatusOK, customer), appErr)
}
