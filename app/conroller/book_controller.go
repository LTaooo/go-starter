package controller

import (
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

func (b *BookController) Register(engine *gin.Engine) {
	engine.GET("/book", b.GetBook)
}

func (b *BookController) GetBook(c *gin.Context) {
	response.NewResponse().Success(c, "book")
}
