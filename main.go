package main

import (
	"log"

	"github.com/minhdung/nailstore/internal/controller"
	db "github.com/minhdung/nailstore/internal/infrastructure/db"
	"github.com/minhdung/nailstore/internal/infrastructure/repositories"
	"github.com/minhdung/nailstore/internal/usecase"
)

const (
	dbSource      = "root:020920@tcp(localhost:3306)/nailstore?charset=utf8mb4&parseTime=True&loc=Local"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := db.InitDB(dbSource)
	if err != nil {
		log.Fatal("can not donnect to db:", err)
	}
	userRepo := repositories.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	accountController := controller.NewAccountController(userUsecase)
	server := controller.NewServer(accountController)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start a server:", err)
	}
}
