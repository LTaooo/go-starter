package server

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
	"go.uber.org/fx"
)

func setupRouter(routeHandler *route.RouteHandler) *gin.Engine {
	// 1. 创建 gin 引擎，不使用默认中间件
	engine := gin.New()

	// 2. 使用自定义日志中间件
	engine.Use(middleware.GinLogger())
	engine.Use(middleware.GinRecovery())
	engine.Use(middleware.ErrorHandler())

	// 3. 通过依赖注入的路由处理器初始化路由
	routeHandler.Init(engine)

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
		logger.SugaredLogger.Error("Mysql连接失败", "error", err)
		panic(err)
	}
	logger.SugaredLogger.Info("Mysql连接成功")
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

func Init(routeHandler *route.RouteHandler) *gin.Engine {
	// 1. 初始化日志系统
	initLogger()

	// 2. 加载配置
	config.LoadConfig()

	// 3. 初始化 Swagger 文档
	initSwagger()

	// 4. 设置 Gin 模式
	setGinMode()

	// 5. 初始化数据库连接
	initDatabase()

	// 6. 初始化 Redis 连接
	initRedis()

	// 7. 通过依赖注入设置路由
	return setupRouter(routeHandler)
}

func NewHTTPServer(lc fx.Lifecycle, routeHandler *route.RouteHandler) *http.Server {
	// 1. 通过依赖注入初始化应用
	r := Init(routeHandler)

	// 2. 创建 HTTP 服务器
	server := &http.Server{
		Addr:    config.GetConfig().GetListenAddr(),
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.SugaredLogger.Info("项目启动成功:", config.GetConfig().GetListenAddr(), "+", config.GetConfig().AppEnv)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.SugaredLogger.Error("服务器启动失败", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			gracefulShutdown(server)
			return nil
		},
	})
	return server
}
