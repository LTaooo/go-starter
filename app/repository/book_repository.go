package repository

import (
	"go-starter/app/model"
	"go-starter/core/database"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		db: database.MySQL,
	}
}

func (r *BookRepository) Get(id uint) (*model.Book, error) {
	var book model.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) Create(book *model.Book) error {
	result := r.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
