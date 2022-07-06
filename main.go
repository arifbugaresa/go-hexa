package main

import (
	"fmt"
	"github.com/arifbugaresa/go-hexa/api/v1"
	absensiContr "github.com/arifbugaresa/go-hexa/api/v1/absensi"
	aktivitasContr "github.com/arifbugaresa/go-hexa/api/v1/aktivitas"
	"github.com/arifbugaresa/go-hexa/api/v1/health"
	userInfoContr "github.com/arifbugaresa/go-hexa/api/v1/user_info"
	absensiServ "github.com/arifbugaresa/go-hexa/business/absensi"
	aktivitasServ "github.com/arifbugaresa/go-hexa/business/aktivitas"
	userInfoServ "github.com/arifbugaresa/go-hexa/business/user_info"
	configuration "github.com/arifbugaresa/go-hexa/config"
	"github.com/arifbugaresa/go-hexa/middleware/auth"
	absensiRepo "github.com/arifbugaresa/go-hexa/modules/absensi"
	aktivitasRepo "github.com/arifbugaresa/go-hexa/modules/aktivitas"
	"github.com/arifbugaresa/go-hexa/modules/database"
	userInfoRepo "github.com/arifbugaresa/go-hexa/modules/user_info"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strconv"
)

var (
	config       = configuration.GetConfig()
	dbConnection = database.NewDatabaseConnection(config)
	e            = echo.New()
)

var (
	authService         = auth.NewService()
	userInfoRepository  = userInfoRepo.NewRepository(dbConnection)
	userInfoService     = userInfoServ.NewService(userInfoRepository)
	userInfoController  = userInfoContr.NewController(userInfoService, authService)
	HealthController    = health.NewController()
	absensiRepository   = absensiRepo.NewRepository(dbConnection)
	absensiSevice       = absensiServ.NewService(absensiRepository)
	absensiController   = absensiContr.NewController(absensiSevice)
	aktivitasRepository = aktivitasRepo.NewRepository(dbConnection)
	aktivitasService    = aktivitasServ.NewService(aktivitasRepository, absensiRepository)
	aktivitasController = aktivitasContr.NewController(aktivitasService)
)

func main() {
	defer database.CloseDatabaseConnection(dbConnection)

	migrateDatabase()

	api.Controller(e, HealthController, userInfoController, absensiController, aktivitasController)

	runServer()
}

func migrateDatabase() {
	dbConnection.AutoMigrate(
		&userInfoServ.UserInfo{},
		&absensiServ.Absensi{},
		&aktivitasServ.Aktivitas{},
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
