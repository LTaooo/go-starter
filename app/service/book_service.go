package service

import (
	"errors"
	"go-starter/app/model"
	"go-starter/core/database"
	"go-starter/core/logger"

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
	// 1. 参数验证
	if id == 0 {
		return nil, errors.New("书籍ID不能为空")
	}

	// 2. 查询数据库
	var book model.Book
	result := s.db.First(&book, id)

	// 3. 处理查询结果
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.SugaredLogger.Warn("书籍不存在", "id", id)
			return nil, errors.New("书籍不存在")
		}
		logger.SugaredLogger.Error("查询书籍失败", "id", id, "error", result.Error)
		return nil, result.Error
	}

	logger.SugaredLogger.Info("查询书籍成功", "id", id, "name", book.Name)
	return &book, nil
}
