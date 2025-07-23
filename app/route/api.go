package route

import (
	"go-starter/app/config"
	controller "go-starter/app/conroller"
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

/**
 * 初始化路由
 */
func Init(engine *gin.Engine) {
	initDefaultRoutes(engine)
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
		response.NewResponse().Success(c, config.GetConfig().AppName)
	})

	engine.NoRoute(func(c *gin.Context) {
		response.NewResponse().Error(c, enum.HttpNotFound, "Not Found")
	})

	engine.NoMethod(func(c *gin.Context) {
		response.NewResponse().Error(c, enum.BadRequest, "Method Not Allowed")
	})
}
