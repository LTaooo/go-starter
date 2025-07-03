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
	err := logger.Init(logger.NewConfig())
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Get().Debug("这是一条调试级别的日志")
}

func main() {
	initLogger()
	r := setupRouter()
	config.LoadConfig()
	r.Run(":8080")
}
