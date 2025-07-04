package logger

import (
	"github.com/gin-gonic/gin"
)

// GinLogger 返回一个 gin.HandlerFunc 中间件，使用 Zap 记录日志
func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 1. 记录请求日志
		SugaredLogger.Info(
			"Method: ", param.Method,
			" Path: ", param.Path,
			" Status: ", param.StatusCode,
			" ClientIP: ", param.ClientIP,
		)

		return ""
	})
}

// GinRecovery 返回一个 gin.HandlerFunc 中间件，使用 Zap 记录 panic 恢复
func GinRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// 1. 记录 panic 信息
		SugaredLogger.Error("HTTP 请求发生 panic",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"client_ip", c.ClientIP(),
			"panic", recovered,
		)

		// 2. 返回 500 错误
		c.AbortWithStatus(500)
	})
}
