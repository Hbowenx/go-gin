package request

type ProductListRequest struct {
	LabelId  string `form:"label_id" binding:"required"`
	Desc string `form:"desc"`
	//Page int `form:"page" binding:"required"`
}

type ProductAddRequest struct {
	Desc    string `form:"desc" binding:"required"`
	Images      []string `form:"images" binding:"required"`
	Sort    int `form:"sort" binding:"required"`
	LabelId int `form:"label_id" binding:"required"`
}

type ProductInfoRequest struct {
	Id int `form:"id" binding:"required"`
}

type ProductDelRequest struct {
	Id int `form:"id" binding:"required"`
}