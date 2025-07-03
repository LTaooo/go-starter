package route

import (
	controller "go-starter/app/conroller"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	controller.NewBookController().Register(engine)
}
