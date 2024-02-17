package domain

import (
	"banking/customErrors"
	"banking/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *customErrors.AppError) {
	saveAccountQuery := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?)"

	result, err := d.client.Exec(saveAccountQuery, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error(fmt.Sprintf("Error while creating a new account: %v", err.Error()))

		return nil, customErrors.NewInternalServerError("failed to save account")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error(fmt.Sprintf("Error while getting lats insert id for a new account: %v", err.Error()))

		return nil, customErrors.NewInternalServerError("failed to save account")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
