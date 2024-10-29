package main

import (
	"log"

	api "github.com/minhdung/nailstore/internal/api"
	"github.com/minhdung/nailstore/internal/api/handler"
	db "github.com/minhdung/nailstore/internal/infrastructure/db"
	"github.com/minhdung/nailstore/internal/infrastructure/repositories"
	usecase "github.com/minhdung/nailstore/internal/usecase"
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
	accountHandler := CreateAccountHandler(config, conn)
	userHandler := CreateUserHandler(config, conn)

	server, nil := api.NewServer(config, accountHandler, userHandler)
	if err != nil {
		log.Fatal("can not create a server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start a server:", err)
	}
}

func CreateAccountHandler(util util.Config, conn *gorm.DB) *handler.AccountHandler {
	userRepo := repositories.NewUserRepository(conn)
	sessionRepo := repositories.NewSessionRepository(conn)
	userUsecase := usecase.NewUserUsecaseImpl(sessionRepo, userRepo)
	accountController, err := handler.NewAccountHandler(util, userUsecase)
	if err != nil {
		log.Fatal("can not create account handler:", err)
	}
	return accountController
}

func CreateUserHandler(util util.Config, conn *gorm.DB) *handler.UserHandler {
	userRepo := repositories.NewUserRepository(conn)
	sessionRepo := repositories.NewSessionRepository(conn)
	userUsecase := usecase.NewUserUsecaseImpl(sessionRepo, userRepo)
	accountController, err := handler.NewUserHandler(util, userUsecase)
	if err != nil {
		log.Fatal("can not create account handler:", err)
	}
	return accountController
}
