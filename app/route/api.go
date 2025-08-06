package route

import (
	controller "go-starter/app/conroller"
	"go-starter/core/config"
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouteHandler 路由处理器结构体
type RouteHandler struct {
	bookController *controller.BookController
}

// NewRouteHandler 创建路由处理器
func NewRouteHandler(bookController *controller.BookController) *RouteHandler {
	return &RouteHandler{
		bookController: bookController,
	}
}

/**
 * 初始化路由
 */
func (r *RouteHandler) Init(engine *gin.Engine) {
	// 1. 初始化默认路由
	initDefaultRoutes(engine)

	// 2. 添加swagger路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 3. 创建API路由组
	api := engine.Group("/api")

	r.registerBookRoutes(api)
}

/**
 * 注册书籍相关路由
 */
func (r *RouteHandler) registerBookRoutes(api *gin.RouterGroup) {
	api.GET("/book", r.bookController.GetBook)
	api.POST("/book/create", r.bookController.CreateBook)
}

/**
 * 初始化默认路由
 */
func initDefaultRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		response.Success(c, config.GetConfig().AppName)
	})

	engine.NoRoute(func(c *gin.Context) {
		response.Error(c, enum.HttpNotFound, "Not Found")
	})

	engine.NoMethod(func(c *gin.Context) {
		response.Error(c, enum.BadRequest, "Method Not Allowed")
	})
}
