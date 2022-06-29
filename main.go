package main

import (
	"fmt"
	"github.com/arifbugaresa/go-hexa/api/v1"
	"github.com/arifbugaresa/go-hexa/business/user_info"
	configuration "github.com/arifbugaresa/go-hexa/config"
	"github.com/arifbugaresa/go-hexa/modules/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strconv"
)

var (
	config       = configuration.GetConfig()
	dbConnection = database.NewDatabaseConnection(config)
	e            = echo.New()
)

func main()  {
	defer database.CloseDatabaseConnection(dbConnection)

	migrateDatabase()

	api.Controller(e)

	runServer()
}

func migrateDatabase() {
	dbConnection.AutoMigrate(
		&user_info.UserInfo{},
	)

	log.Info("Success migrate database, " + strconv.Itoa(int(dbConnection.RowsAffected)) + " row affected.")
}

func runServer() {
	address := fmt.Sprintf("localhost:%s", config.AppPort)
	err := e.Start(address)
	if err != nil {
		log.Info("shutting down the server")
	}
}

