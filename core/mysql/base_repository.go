package database

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Find(id uint) (*T, error) {
	var entity T
	result := r.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
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
