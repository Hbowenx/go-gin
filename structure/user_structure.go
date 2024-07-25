package structure

type GetUserInfo struct {
	Openid  string `gorm:"column:openid"`
}

type UserLogin struct {
	Openid  string `gorm:"column:openid"`
	Phone string `gorm:"column:phone"`
	NickName string `gorm:"column:nick_name"`
	AvatarUrl string `gorm:"column:avatar_url"`
}

type UserInfo struct {
	Id        int         `json:"id"`
	Openid string      `json:"openid"`
	Admin int `json:"admin"`
	Phone     string      `json:"phone"`
	NickName      string      `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}