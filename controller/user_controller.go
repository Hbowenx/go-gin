package controller

import (
	"cell_phone_store/model"
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func UserLogin(c *gin.Context)  {
	var userLoginRequest request.UserLoginRequest
	if err := c.ShouldBindWith(&userLoginRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	userLogin := structure.UserLogin{
		Openid: userLoginRequest.Openid,
		Phone: userLoginRequest.Phone,
		NickName: userLoginRequest.NickName,
		AvatarUrl:userLoginRequest.AvatarUrl,
	}
	userInfo ,err :=model.UserLogin(userLogin)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, userInfo)
}

func UserInfo(c *gin.Context)   {
	var userInfoRequest request.UserInfoRequest
	if err := c.ShouldBindWith(&userInfoRequest,binding.Form); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	getUserInfo := structure.GetUserInfo{
		Openid: userInfoRequest.Openid,
	}
	userInfo ,err :=model.UserInfo(getUserInfo)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, userInfo)
}