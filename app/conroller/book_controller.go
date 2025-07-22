package controller

import (
	"strconv"

	"go-starter/app/service"
	"go-starter/core/enum"
	"go-starter/core/response"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		bookService: service.NewBookService(),
	}
}

func (b *BookController) Register(engine *gin.Engine) {
	engine.GET("/book/:id", b.GetBook)
}

func (b *BookController) GetBook(c *gin.Context) {
	// 1. 获取路径参数中的ID
	idStr := c.Param("id")
	if idStr == "" {
		response.NewResponse().Error(c, enum.BadRequest, "书籍ID不能为空")
		return
	}

	// 2. 转换ID为uint类型
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.NewResponse().Error(c, enum.BadRequest, "书籍ID格式错误")
		return
	}

	// 3. 调用服务层查询书籍
	book, err := b.bookService.GetBookByID(uint(id))
	if err != nil {
		response.NewResponse().Error(c, enum.NotFound, err.Error())
		return
	}

	// 4. 返回成功响应
	response.NewResponse().Success(c, book)
}
