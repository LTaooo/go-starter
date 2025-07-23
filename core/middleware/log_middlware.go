package middleware

import (
	"go-starter/core/logger"

	"github.com/gin-gonic/gin"
)

// GinLogger 返回一个 gin.HandlerFunc 中间件，使用 Zap 记录日志
func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 1. 记录请求日志
		logger.SugaredLogger.Info(
			"Method: ", param.Method,
			" Path: ", param.Path,
			" Status: ", param.StatusCode,
			" ClientIP: ", param.ClientIP,
		)

		return ""
	})
}
