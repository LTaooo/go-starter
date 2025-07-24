package http

import (
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

func FromQuery(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindQuery(req); err != nil {
		response.Error(c, enum.BadRequest, err.Error())
		return err
	}
	return nil
}

func FromJson(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		response.Error(c, enum.BadRequest, err.Error())
		return err
	}
	return nil
}
