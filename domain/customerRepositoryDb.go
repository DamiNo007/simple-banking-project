package domain

import (
	"banking/customErrors"
	"banking/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *customErrors.AppError) {
	findAllQuery := "select * from customers"
	customers := make([]Customer, 0)

	err := d.client.Select(&customers, findAllQuery)

	if err != nil {
		logger.Error(fmt.Sprintf("Error while scanning customers by status %v", err.Error()))
		return nil, customErrors.NewInternalServerError("failed to retrieve customers")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindByStatus(status string) ([]Customer, *customErrors.AppError) {
	findByStatusQuery := "select * from customers where status = ?"
	customers := make([]Customer, 0)

	err := d.client.Select(&customers, findByStatusQuery, status)

	if err != nil {
		logger.Error(fmt.Sprintf("Error while scanning customers by status %v", err.Error()))
		return nil, customErrors.NewInternalServerError("failed to retrieve customers by status")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *customErrors.AppError) {
	findByIdQuery := "select * from customers where customer_id = ?"
	var c Customer

	err := d.client.Get(&c, findByIdQuery, id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			logger.Error(fmt.Sprintf("Customer not found for id %v", id))
			return nil, customErrors.NewNotFoundError("customer not found")
		default:
			logger.Error(fmt.Sprintf("Error while scanning customer by id %v", err.Error()))
			return nil, customErrors.NewInternalServerError("failed to retrieve customer by id")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
