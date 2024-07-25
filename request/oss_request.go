package request

type UploadOssRequest struct {
	Dir string `form:"dir" binding:"required"`
}
