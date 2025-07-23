package repository

import (
	"go-starter/app/model"
	"go-starter/core/database"
)

type BookRepository struct {
	database.BaseRepository[model.Book]
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		BaseRepository: *database.NewBaseRepository[model.Book](database.MySQL),
	}
}
