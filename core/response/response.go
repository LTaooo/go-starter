package response

import (
	"go-starter/core/enum"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    enum.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	rs := Response{
		Code:    enum.OK,
		Message: enum.OK.Message(),
		Data:    data,
	}
	ctx.JSON(200, rs)
}

func Error(ctx *gin.Context, code enum.Code, message string) {
	rs := Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	ctx.JSON(200, rs)
}
