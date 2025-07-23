package service

import (
	dto "go-starter/app/dto/book"
	"go-starter/app/model"
	"go-starter/core/database"
	"go-starter/core/utils/datetime"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService() *BookService {
	return &BookService{
		db: database.MySQL,
	}
}

// GetBookByID 通过ID查询书籍
func (s *BookService) GetBookByID(id uint) (*model.Book, error) {
	var book model.Book
	result := s.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (s *BookService) CreateBook(req dto.BookCreateReq) (*model.Book, error) {
	var book model.Book
	book.Name = req.Name
	book.Author = req.Author
	book.Price = req.Price
	book.CreatedAt = datetime.FromNow().Timestamp()
	book.UpdatedAt = datetime.FromNow().Timestamp()
	result := s.db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}
