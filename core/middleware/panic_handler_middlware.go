package middleware

import (
	"go-starter/core/enum"
	"go-starter/core/logger"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

func GinRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.SugaredLogger.Error("HTTP 请求发生 panic",
			"Method: ", c.Request.Method,
			" Path: ", c.Request.URL.Path,
			" ClientIP: ", c.ClientIP(),
			" Panic: ", recovered,
		)

		response.Error(c, enum.InternalError, "系统异常")
	})
}
