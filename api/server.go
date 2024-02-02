package api

import (
	db "github.com/Mive667/GoBankHub/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for backend API.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to the router
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
