package domain

import (
	"banking/customErrors"
	"banking/dto"
)

type CustomerList []Customer

type Customer struct {
	Id          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"fullName" xml:"fullName"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *customErrors.AppError)
	FindById(string) (*Customer, *customErrors.AppError)
	FindByStatus(status string) ([]Customer, *customErrors.AppError)
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "not_active"
	}

	return statusAsText
}

func (c *Customer) ToDTO() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

func (cl CustomerList) ToDTO() []dto.CustomerResponse {
	customerResponseList := make([]dto.CustomerResponse, len(cl))

	for id, c := range cl {
		customerResponseList[id] = *c.ToDTO()
	}

	return customerResponseList
}
