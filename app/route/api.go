package route

import (
	"go-starter/app/config"
	controller "go-starter/app/conroller"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

/**
 * 初始化路由
 */
func Init(engine *gin.Engine) {
	initDefaultRoutes(engine)
	controller.NewBookController().Register(engine)
}

/**
 * 初始化默认路由
 */
func initDefaultRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		response.NewResponse().Success(c, config.GetConfig().AppName)
	})

	engine.NoRoute(func(c *gin.Context) {
		response.NewResponse().Error(c, response.NotFound, "Not Found")
	})

	engine.NoMethod(func(c *gin.Context) {
		response.NewResponse().Error(c, response.BadRequest, "Method Not Allowed")
	})
}
