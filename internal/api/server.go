package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhdung/nailstore/internal/api/middleware"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	token "github.com/minhdung/nailstore/internal/usecase/token"
	"github.com/minhdung/nailstore/internal/util"
)

// Server serves HTTP request
type Server struct {
	config            util.Config
	router            *gin.Engine
	tokenMaker        interfaceObject.Maker
	accountController *AccountController
}

func NewServer(config util.Config, accountController *AccountController) (*Server, error) {
	token, err := token.NewJWTMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot create token marker : %w", &err)
	}
	server := &Server{
		config:            config,
		accountController: accountController,
		tokenMaker:        token,
	}
	router := gin.Default()

	router.POST("/accounts", server.accountController.CreateAccount)
	router.GET("/accounts/:id", server.accountController.GetAccount)
	// add routers for route
	server.router = router
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.accountController.CreateAccount)
	router.GET("/accounts/:id", server.accountController.GetAccount)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
