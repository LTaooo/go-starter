package service

import (
	dto "go-starter/app/dto/book"
	"go-starter/app/model"
	"go-starter/app/repository"
	"go-starter/core/utils/datetime"
)

type BookService struct {
	bookRepository *repository.BookRepository
}

func NewBookService() *BookService {
	return &BookService{
		bookRepository: repository.NewBookRepository(),
	}
}

// GetBookByID 通过ID查询书籍
func (s *BookService) GetBookByID(id uint) *model.Book {
	return s.bookRepository.Find(id)
}

func (s *BookService) CreateBook(req dto.BookCreateReq) (*model.Book, error) {
	var book model.Book
	book.Name = req.Name
	book.Author = req.Author
	book.Price = req.Price
	book.CreatedAt = datetime.FromNow().Timestamp()
	book.UpdatedAt = datetime.FromNow().Timestamp()
	err := s.bookRepository.Create(&book)
	return &book, err
}
