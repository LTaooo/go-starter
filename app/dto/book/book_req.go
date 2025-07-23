package dto

type BookGetReq struct {
	Id uint `form:"id" binding:"required,min=1"`
}

type BookCreateReq struct {
	// 书籍名称
	Name string `json:"name" binding:"required,min=1,max=255" example:"Go语言编程"`
	// 作者
	Author string `json:"author" binding:"required,min=1,max=255" example:"张三"`
	// 价格
	Price float64 `json:"price" binding:"required,min=0.01" example:"100.00"`
}
