package request

type UserLoginRequest struct {
	Openid  string `form:"openid" binding:"required"`
	NickName string `form:"nick_name" binding:"required"`
	Phone string `form:"phone" binding:"required"`
	AvatarUrl string `form:"avatar_url" binding:"required"`
}

type UserInfoRequest struct {
	Openid  string `form:"openid" binding:"required"`
}