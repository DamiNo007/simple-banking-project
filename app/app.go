package app

import (
	"banking/apiServer"
	"banking/apiServer/handlers"
	"banking/app/config"
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func Start() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dbConn, err := startMySqlDbConn(cfg)

	if err != nil {
		panic(err)
	}

	//creating handlers
	//customerHandlers := handlers.NewCustomerHandler(service.NewCustomerService(domain.NewCustomerRepositoryStub()))
	customerHandlers := handlers.NewCustomerHandler(service.NewCustomerService(domain.NewCustomerRepositoryDb(dbConn)))
	accountHandlers := handlers.NewAccountHandler(service.NewAccountService(domain.NewAccountRepositoryDb(dbConn)))

	//creating a server
	server := apiServer.NewApiServer(
		cfg.ApiServer.Host,
		cfg.ApiServer.Port,
		&customerHandlers,
		&accountHandlers,
	)

	//starting the server
	log.Fatal(server.Start())
}

func startMySqlDbConn(cfg *config.Config) (*sqlx.DB, error) {
	dbHost := cfg.MySql.Host
	dbPort := cfg.MySql.Port
	dbUserName := cfg.MySql.UserName
	dbPassword := cfg.MySql.Password
	dbName := cfg.MySql.DbName

	dbClient, err := sqlx.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUserName, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)

	return dbClient, nil
}
