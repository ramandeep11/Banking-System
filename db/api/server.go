package api

import (
	"fmt"
	db "simplebank/db/sqlc"
	"simplebank/db/token"
	"simplebank/db/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(confi util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewpasetoMaker(confi.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}

	server := &Server{
		config:     confi,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/user", server.createUser)
	router.POST("/users/login",server.loginUser)

	authroutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authroutes.POST("/accounts", server.createAccount)
	authroutes.GET("/accounts/:id", server.getAccount)
	authroutes.GET("/accounts", server.listAccount)
	authroutes.POST("/transfer", server.createTransfer)

	server.router = router
}

// to start listening for api requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error}
}
