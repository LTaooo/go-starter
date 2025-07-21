package model

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}
