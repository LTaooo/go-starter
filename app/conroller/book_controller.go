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

// @Summary 获取书籍
// @Description 根据ID获取书籍信息
// @Tags book
// @Accept json
// @Produce json
// @Param id query int true "书籍ID" minimum(1)
// @Success 200 {object} response.Response{data=dto.BookGetRes}
// @Router /api/book [get]
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

// @Summary 创建书籍
// @Description 创建新的书籍
// @Tags book
// @Accept json
// @Produce json
// @Param book body dto.BookCreateReq true "书籍信息"
// @Success 200 {object} response.Response{data=dto.BookCreateRes}
// @Router /api/book/create [post]
func (b *BookController) CreateBook(c *gin.Context) {
	var req dto.BookCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewResponse().Error(c, enum.BadRequest, err.Error())
		return
	}

	book, err := b.bookService.CreateBook(req)
	if err != nil {
		response.NewResponse().Error(c, enum.InternalError, err.Error())
		return
	}

	response.NewResponse().Success(c, dto.BookCreateRes{
		Id: book.ID,
	})
}
