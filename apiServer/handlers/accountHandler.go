package handlers

import (
	"banking/apiServer/wrappers"
	"banking/customErrors"
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"
)

type AccountsHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) AccountsHandler {
	return AccountsHandler{service: service}
}

func (h *AccountsHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	format := retrieveFormatFromRequestHeaders(r)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handleApiResponse(w, format, nil, customErrors.NewBadRequestError("failed to decode the request body"))
	} else {
		account, err := h.service.NewAccount(request)
		handleApiResponse(w, format, wrappers.NewResponseWrapper(http.StatusCreated, account), err)
	}
}
