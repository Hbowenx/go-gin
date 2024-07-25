package model

import (
	"cell_phone_store/structure"
	"errors"
)

var WxUserTable = "wx_users"

func WxLogin(wxUserAdd structure.WxUserAdd) error {
	var wxUserRecord struct{
		Id        int
	}
	DB.Table(WxUserTable).Select("id").Where("openid = ?",wxUserAdd.Openid).First(&wxUserRecord)
	if wxUserRecord.Id == 0 {
		DB.Table(WxUserTable).Create(&wxUserAdd)
	}
	return nil
}

func UserLogin(userLogin structure.UserLogin) (structure.UserInfo,error)  {
	var wxUserIdRecord struct{
		Id        int
	}
	DB.Table(WxUserTable).Select("id").Where("openid = ?",userLogin.Openid).First(&wxUserIdRecord)
	if wxUserIdRecord.Id == 0 {
		return structure.UserInfo{},errors.New("未登录,请重新打开小程序。")
	}
	var wxUserInfoRecord struct{
		Id int
		SessionKey string
		Openid string
		UnionId string
		NickName string
		AvatarUrl string
		Phone string
		Admin int
		CreatedAt string
		UpdatedAt string
	}

	DB.Table(WxUserTable).Where("phone = ?",userLogin.Phone).First(&wxUserInfoRecord)
	if wxUserInfoRecord.Id != 0 && wxUserInfoRecord.Openid != userLogin.Openid{
		return structure.UserInfo{},errors.New("手机号已存在！")
	}
	if wxUserInfoRecord.Phone == "" {
		wxUserInfoUpdate :=structure.WxUserInfoUpdate{
			NickName: userLogin.NickName,
			Phone: userLogin.Phone,
			AvatarUrl: userLogin.AvatarUrl,
		}
		DB.Table(WxUserTable).Where("id = ?",wxUserIdRecord.Id).Updates(&wxUserInfoUpdate)
	}

	DB.Table(WxUserTable).Where("id = ?",wxUserIdRecord.Id).First(&wxUserInfoRecord)
	return structure.UserInfo{
		Id: wxUserInfoRecord.Id,
		AvatarUrl:wxUserInfoRecord.AvatarUrl,
		Phone:wxUserInfoRecord.Phone,
		NickName:wxUserInfoRecord.NickName,
		Admin:wxUserInfoRecord.Admin,
		Openid:wxUserInfoRecord.Openid,
		CreatedAt:wxUserInfoRecord.CreatedAt,
		UpdatedAt:wxUserInfoRecord.UpdatedAt,
	},nil
}

func UserInfo(getUserInfo structure.GetUserInfo) (structure.UserInfo, error) {
	var wxUserInfoRecord struct{
		Id int
		SessionKey string
		Openid string
		UnionId string
		NickName string
		AvatarUrl string
		Phone string
		Admin int
		CreatedAt string
		UpdatedAt string
	}
	DB.Table(WxUserTable).Where("openid = ?",getUserInfo.Openid).First(&wxUserInfoRecord)
	if wxUserInfoRecord.Id == 0 {
		return structure.UserInfo{},errors.New("用户信息不存在！")
	}
	return structure.UserInfo{
		Id: wxUserInfoRecord.Id,
		AvatarUrl:wxUserInfoRecord.AvatarUrl,
		Phone:wxUserInfoRecord.Phone,
		NickName:wxUserInfoRecord.NickName,
		Admin:wxUserInfoRecord.Admin,
		Openid:wxUserInfoRecord.Openid,
		CreatedAt:wxUserInfoRecord.CreatedAt,
		UpdatedAt:wxUserInfoRecord.UpdatedAt,
	},nil
}