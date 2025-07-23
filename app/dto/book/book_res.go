package dto

type BookGetRes struct {
	// 书籍ID
	Id uint `json:"id"`
	// 书籍名称
	Name string `json:"name"`
	// 作者
	Author string `json:"author"`
	// 价格
	Price float64 `json:"price"`
	// 创建时间
	CreateAt string `json:"create_at"`
}

type BookCreateRes struct {
	// 书籍id
	Id uint `json:"id"`
}
