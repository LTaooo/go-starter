package middleware

import (
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			response.NewResponse().Error(c, enum.InternalError, err.Error())
		}
	}
}
