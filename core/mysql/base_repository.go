package database

import (
	"errors"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Find(id uint) *T {
	var entity T
	result := r.db.First(&entity, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	if result.Error != nil {
		panic(result.Error)
	}
	return &entity
}

func (r *BaseRepository[T]) Create(entity *T) error {
	result := r.db.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BaseRepository[T]) Update(entity *T) error {
	result := r.db.Save(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	result := r.db.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
