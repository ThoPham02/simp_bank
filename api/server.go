package api

import (
	db "github.com/ThoPham02/simp_bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	Store  db.Store
	Router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{Store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", ValidCurrency)
	}

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.getListAccounts)
	router.PUT("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	router.POST("/transfer", server.createTransfer)

	server.Router = router
	return server
}

func errorResponse(err error) gin.H {
	customErr := HandleValidatorError(err)

	return gin.H{"error": customErr}
}
