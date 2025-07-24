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

/**
 * 初始化路由
 */
func Init(engine *gin.Engine) {
	initDefaultRoutes(engine)
	// 添加swagger路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := engine.Group("/api")

	// 书籍路由
	bookController := controller.NewBookController()
	api.GET("/book", bookController.GetBook)
	api.POST("/book/create", bookController.CreateBook)
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
