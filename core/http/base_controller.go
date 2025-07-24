package http

import (
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (t *BaseController) IsError(c *gin.Context, err error) bool {
	if err != nil {
		response.Error(c, enum.InternalError, err.Error())
		return true
	}
	return false
}
