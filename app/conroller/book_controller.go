package controller

import (
	dto "go-starter/app/dto/book"
	"go-starter/app/service"
	"go-starter/core/enum"
	"go-starter/core/response"
	"go-starter/core/utils/datetime"

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
	engine.GET("/book", b.GetBook)
}

func (b *BookController) GetBook(c *gin.Context) {
	var req dto.BookGetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewResponse().Error(c, enum.BadRequest, err.Error())
		return
	}

	book, err := b.bookService.GetBookByID(req.Id)

	if err != nil {
		response.NewResponse().Success(c, nil)
		return
	}

	response.NewResponse().Success(c, dto.BookGetRes{
		Id:       book.ID,
		Name:     book.Name,
		Author:   book.Author,
		Price:    book.Price,
		CreateAt: datetime.FromTimestamp(book.CreatedAt).Datetime(),
	})
}
