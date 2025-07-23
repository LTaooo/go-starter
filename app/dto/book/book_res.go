package dto

type BookGetRes struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	CreateAt string  `json:"create_at"`
}

type BookCreateRes struct {
	Id uint `json:"id"`
}
