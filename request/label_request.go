package request

type LabelAddRequest struct {
	Name string `form:"name" binding:"required"`
	Sort   int `form:"sort" binding:"required"`
}

type LabelDelRequest struct {
	Id int `form:"id" binding:"required"`
}