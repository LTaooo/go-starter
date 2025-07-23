package dto

type BookGetReq struct {
	Id uint `form:"id" binding:"required"`
}

type BookGetRes struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	CreateAt string  `json:"create_at"`
}
