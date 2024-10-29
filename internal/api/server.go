package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhdung/nailstore/internal/api/handler"
	"github.com/minhdung/nailstore/internal/api/middleware"
	interfaceObject "github.com/minhdung/nailstore/internal/interface"
	token "github.com/minhdung/nailstore/internal/usecase/token"
	"github.com/minhdung/nailstore/internal/util"
)

// Server serves HTTP request
type Server struct {
	config         util.Config
	router         *gin.Engine
	tokenMaker     interfaceObject.Maker
	accountHandler *handler.AccountHandler
	userHandler    *handler.UserHandler
}

func NewServer(config util.Config,
	accountController *handler.AccountHandler,
	userHandler *handler.UserHandler) (*Server, error) {
	token, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token marker : %w", &err)
	}
	server := &Server{
		config:         config,
		accountHandler: accountController,
		userHandler:    userHandler,
		tokenMaker:     token,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// Set up the router once
	router := gin.Default()

	// Apply auth middleware to specific routes
	authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.accountHandler.CreateAccount)

	// Unauthenticated routes
	router.POST("/user", server.userHandler.CreateUser)
	router.POST("/login", server.accountHandler.LoginUser)
	router.GET("/accounts/:id", server.accountHandler.GetAccount)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
