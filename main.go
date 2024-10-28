package main

import (
	"log"

	api "github.com/minhdung/nailstore/internal/api"
	db "github.com/minhdung/nailstore/internal/infrastructure/db"
	"github.com/minhdung/nailstore/internal/infrastructure/repositories"
	usecase "github.com/minhdung/nailstore/internal/usecase/user"
	"github.com/minhdung/nailstore/internal/util"
	"gorm.io/gorm"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	conn, err := db.InitDB(config.DbSource)
	if err != nil {
		log.Fatal("can not donnect to db:", err)
	}
	accountController := CreateAccountController(conn)
	server, nil := api.NewServer(config, accountController)
	if err != nil {
		log.Fatal("can not create a server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start a server:", err)
	}
}

func CreateAccountController(conn *gorm.DB) *api.AccountController {
	userRepo := repositories.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	accountController := api.NewAccountController(userUsecase)
	return accountController
}
