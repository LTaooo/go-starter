package model

type Book struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string  `json:"name" gorm:"type:varchar(255);not null"`
	Author    string  `json:"author" gorm:"type:varchar(255)"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime"`
}
