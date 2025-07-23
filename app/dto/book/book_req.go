package dto

type BookGetReq struct {
	Id uint `form:"id" binding:"required,min=1"`
}

type BookCreateReq struct {
	Name   string  `json:"name" binding:"required,min=1,max=255"`
	Author string  `json:"author" binding:"required,min=1,max=255"`
	Price  float64 `json:"price" binding:"required,min=0.01"`
}
