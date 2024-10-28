package main

import (
	"log"

	api "github.com/minhdung/nailstore/internal/api"
	db "github.com/minhdung/nailstore/internal/infrastructure/db"
	"github.com/minhdung/nailstore/internal/infrastructure/repositories"
	"github.com/minhdung/nailstore/internal/usecase"
	"github.com/minhdung/nailstore/internal/util"
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
	userRepo := repositories.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	accountController := api.NewAccountController(userUsecase)
	server := api.NewServer(accountController)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start a server:", err)
	}
}
