package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() Response {
	return Response{}
}

func (rs Response) Success(ctx *gin.Context, data interface{}) {
	rs = Response{
		Code:    OK,
		Message: OK.Message(),
		Data:    data,
	}
	ctx.JSON(200, rs)
}

func (rs Response) Error(ctx *gin.Context, code Code, message string) {
	rs = Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	ctx.JSON(200, rs)
}
