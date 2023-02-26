package api

import (
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}

func errorRequestBinding(err error) gin.H {
	return gin.H{"error": HandleValidatorError(err)}
}
