package main

import (
	"go-starter/app/config"
	"go-starter/app/route"
	"go-starter/core/logger"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	engine := gin.Default()
	route.Init(engine)
	return engine
}
func initLogger() {
	// 1. 初始化日志系统
	if err := logger.InitLogger(); err != nil {
		panic(err)
	}

	// 2. 使用 Zap logger
	logger.SugaredLogger.Info("这是一条信息级别的日志")
	logger.SugaredLogger.Error("这是一条错误级别的日志")
}

func main() {
	initLogger()
	r := setupRouter()
	gin.SetMode(gin.DebugMode)
	config.LoadConfig()
	r.Run(":8080")
}
