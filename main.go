package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-starter/app/route"
	"go-starter/core/config"
	"go-starter/core/enum"
	"go-starter/core/logger"
	"go-starter/core/middleware"
	database "go-starter/core/mysql"
	"go-starter/core/redis"

	"go-starter/docs"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// 1. 创建 gin 引擎，不使用默认中间件
	engine := gin.New()

	// 2. 使用自定义日志中间件
	engine.Use(middleware.GinLogger())
	engine.Use(middleware.GinRecovery())
	engine.Use(middleware.ErrorHandler())

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

func initDatabase() {
	// 1. 初始化数据库连接
	if err := database.InitDatabase(); err != nil {
		logger.SugaredLogger.Error("数据库连接失败", "error", err)
		panic(err)
	}
	logger.SugaredLogger.Info("数据库连接成功")
}

func initRedis() {
	// 1. 初始化Redis连接
	if err := redis.InitRedis(); err != nil {
		logger.SugaredLogger.Error("Redis连接失败", "error", err)
		panic(err)
	}
}

func gracefulShutdown(server *http.Server) {
	// 1. 创建信号通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.SugaredLogger.Info("正在关闭服务器...")

	// 2. 创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. 关闭数据库连接
	if err := database.CloseDatabase(); err != nil {
		logger.SugaredLogger.Error("关闭数据库连接失败", "error", err)
	}

	// 4. 关闭Redis连接
	if err := redis.CloseRedis(); err != nil {
		logger.SugaredLogger.Error("关闭Redis连接失败", "error", err)
	}

	// 5. 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		logger.SugaredLogger.Error("服务器关闭失败", "error", err)
	}
}

func initSwagger() {
	conf := config.GetConfig()
	docs.SwaggerInfo.Title = conf.AppName
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Host = conf.GetListenAddr()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	initLogger()
	config.LoadConfig()
	initSwagger()
	setGinMode()
	initDatabase()
	initRedis()
	r := setupRouter()
	server := &http.Server{
		Addr:    config.GetConfig().GetListenAddr(),
		Handler: r,
	}
	logger.SugaredLogger.Info("项目启动成功:", config.GetConfig().GetListenAddr(), "+", config.GetConfig().AppEnv)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.SugaredLogger.Error("服务器启动失败", "error", err)
		}
	}()
	gracefulShutdown(server)
}
