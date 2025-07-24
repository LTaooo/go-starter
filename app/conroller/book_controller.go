package controller

import (
	dto "go-starter/app/dto/book"
	"go-starter/app/service"
	"go-starter/core/http"
	"go-starter/core/response"
	"go-starter/core/utils/datetime"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	http.BaseController
	bookService *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		BaseController: *http.NewBaseController(),
		bookService:    service.NewBookService(),
	}
}

// @Summary 获取书籍
// @Description 根据ID获取书籍信息
// @Tags book
// @Param id query int true "书籍ID" minimum(1)
// @Success 200 {object} response.Response{data=dto.BookGetRes}
// @Router /api/book [get]
func (t *BookController) GetBook(c *gin.Context) {
	var req dto.BookGetReq
	if t.IsError(c, http.FromQuery(c, &req)) {
		return
	}

	book := t.bookService.GetBookByID(req.Id)

	if book != nil {
		response.Success(c, dto.BookGetRes{
			Id:       book.ID,
			Name:     book.Name,
			Author:   book.Author,
			Price:    book.Price,
			CreateAt: datetime.FromTimestamp(book.CreatedAt).Datetime(),
		})
		return
	}
	response.Success(c, nil)
}

// @Summary 创建书籍
// @Description 创建新的书籍
// @Tags book
// @Param book body dto.BookCreateReq true "书籍信息"
// @Success 200 {object} response.Response{data=dto.BookCreateRes}
// @Router /api/book/create [post]
func (t *BookController) CreateBook(c *gin.Context) {
	var req dto.BookCreateReq
	if t.IsError(c, http.FromJson(c, &req)) {
		return
	}

	book, err := t.bookService.CreateBook(req)
	if t.IsError(c, err) {
		return
	}

	response.Success(c, dto.BookCreateRes{
		Id: book.ID,
	})
}
