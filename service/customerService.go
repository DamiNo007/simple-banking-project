package service

import (
	"banking/customErrors"
	"banking/domain"
	"banking/dto"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *customErrors.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *customErrors.AppError)
	GetCustomersByStatus(status string) ([]dto.CustomerResponse, *customErrors.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *customErrors.AppError) {
	customers, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	customerList := domain.CustomerList(customers)

	return customerList.ToDTO(), nil
}

func (s DefaultCustomerService) GetCustomersByStatus(status string) ([]dto.CustomerResponse, *customErrors.AppError) {
	customers, err := s.repository.FindByStatus(status)

	if err != nil {
		return nil, err
	}

	customerList := domain.CustomerList(customers)

	return customerList.ToDTO(), nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *customErrors.AppError) {
	customer, err := s.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return customer.ToDTO(), nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
