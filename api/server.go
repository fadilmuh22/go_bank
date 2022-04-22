package api

import (
	"fmt"

	"fadilmuh22/go_bank/db"
	"fadilmuh22/go_bank/token"
	"fadilmuh22/go_bank/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	setupRouter(server)

	return server, nil
}

func setupRouter(server *Server) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {})

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/accounts/:id", server.getAccount)

	server.router = router
}

func (server *Server) RunServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
