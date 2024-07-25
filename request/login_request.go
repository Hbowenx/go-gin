package request

type LoginRequest struct {
	Code string `form:"code" binding:"required"`
}
