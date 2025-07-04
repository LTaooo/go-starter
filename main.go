package main

import (
	"go-starter/app/config"
	"go-starter/app/route"
	"go-starter/core/logger"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// 1. 创建 gin 引擎，不使用默认中间件
	engine := gin.New()

	// 2. 使用我们的自定义日志中间件
	engine.Use(logger.GinLogger())
	engine.Use(logger.GinRecovery())

	// 3. 初始化路由
	route.Init(engine)

	return engine
}
func initLogger() {
	// 1. 初始化日志系统
	if err := logger.InitLogger(); err != nil {
		panic(err)
	}
}

func main() {
	initLogger()
	r := setupRouter()
	gin.SetMode(gin.DebugMode)
	config.LoadConfig()
	logger.SugaredLogger.Info("项目启动成功")
	r.Run(":8080")
}
