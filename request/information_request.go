package request

type InformationListRequest struct {
	Page int `form:"page" binding:"required"`
}

type InformationAddRequest struct {
	MasterMap string `form:"master_map" binding:"required"`
	Title     string `form:"title" binding:"required"`
	Text      string `form:"text" binding:"required"`
	Images      []string `form:"images"`
}

type InformationInfoRequest struct {
	Id int `form:"id" binding:"required"`
}

type InformationDelRequest struct {
	Id int `form:"id" binding:"required"`
}