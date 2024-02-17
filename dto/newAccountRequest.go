package dto

import (
	"banking/customErrors"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customerId" xml:"customerId"`
	AccountType string  `json:"accountType" xml:"accountType"`
	Amount      float64 `json:"amount" xml:"amount"`
}

func (r NewAccountRequest) Validate() *customErrors.AppError {
	if r.Amount < 500 {
		return customErrors.NewStatusUnprocessableEntityError("amount should be greater or equal to 500")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return customErrors.NewStatusUnprocessableEntityError("account type should be \"saving\" or \"checking\"")
	}

	return nil
}
