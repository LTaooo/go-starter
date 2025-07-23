package service

import (
	"go-starter/app/model"
	"go-starter/core/database"

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
