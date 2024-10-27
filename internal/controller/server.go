package controller

import (
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request
type Server struct {
	router            *gin.Engine
	accountController *AccountController
}

func NewServer(accountController *AccountController) *Server {
	server := &Server{accountController: accountController}
	router := gin.Default()

	router.POST("/accounts", server.accountController.CreateAccount)
	router.GET("/accounts/:id", server.accountController.GetAccount)
	// add routers for route
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
