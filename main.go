package main

import (
	"go-starter/app/config"
	"go-starter/app/route"
	"go-starter/core/enum"
	"go-starter/core/logger"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// 1. 创建 gin 引擎，不使用默认中间件
	engine := gin.New()

	// 2. 使用自定义日志中间件
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

func setGinMode() {
	// 1. 根据配置中的环境设置 Gin 模式
	switch config.GetConfig().AppEnv {
	case enum.PROD.String():
		gin.SetMode(gin.ReleaseMode)
	case enum.TEST.String():
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func main() {
	// 1. 初始化日志系统
	initLogger()

	// 2. 加载配置
	config.LoadConfig()

	// 3. 根据配置设置 Gin 运行模式
	setGinMode()

	// 4. 设置路由
	r := setupRouter()

	// 5. 启动服务器
	println(config.GetConfig().AppName)
	logger.SugaredLogger.Info("项目启动成功:", config.GetConfig().GetListenAddr(), "+", config.GetConfig().AppEnv)
	r.Run(config.GetConfig().GetListenAddr())
}
