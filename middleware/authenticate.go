package middleware

import (
	"cell_phone_store/controller"
	"cell_phone_store/model"
	"cell_phone_store/structure"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context)  {
	openid := c.Request.FormValue("openid")
	getUserInfo := structure.GetUserInfo{
		Openid: openid,
	}
	_ ,err :=model.UserInfo(getUserInfo)
	if err != nil {
		controller.HandleError(c, http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.Next()
}

func AuthenticateAdmin(c *gin.Context)  {
	openid := c.Request.FormValue("openid")
	getUserInfo := structure.GetUserInfo{
		Openid: openid,
	}
	userInfo ,err :=model.UserInfo(getUserInfo)
	if err != nil {
		controller.HandleError(c, http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	if userInfo.Admin == 0 {
		controller.HandleError(c, http.StatusInternalServerError, errors.New("不是管理员不能操作"))
		c.Abort()
		return
	}
	c.Next()
}