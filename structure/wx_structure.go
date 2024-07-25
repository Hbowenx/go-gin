package structure

import "time"

type Jscode2SessionResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Openid   string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid  string `json:"unionid,omitempty"`
}

type WxUserAdd struct {
	SessionKey string `gorm:"column:session_key"`
	Openid string `gorm:"column:openid"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type WxUserInfoUpdate struct {
	Phone string `gorm:"column:phone"`
	AvatarUrl string `gorm:"column:avatar_url"`
	NickName string `gorm:"column:nick_name"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}