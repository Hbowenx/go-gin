package controller

import (
	"cell_phone_store/model"
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"net/url"
)

func WxLogin(c *gin.Context)  {
	// 读取配置文件
	var conf structure.Config
	if _, err := toml.DecodeFile(structure.ConfigPath, &conf); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	var loginRequest request.LoginRequest
	if err := c.ShouldBindWith(&loginRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	params := url.Values{}
	params.Add("appid", conf.WxApp.AppId)
	params.Add("secret", conf.WxApp.AppSecret)
	params.Add("js_code", loginRequest.Code)
	params.Add("grant_type", "authorization_code")
	baseURL := "https://api.weixin.qq.com/sns/jscode2session"
	httpBody,err := HttpGet(baseURL,params)

	var result structure.Jscode2SessionResponse
	err = json.Unmarshal(httpBody, &result)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	if result.Errcode !=0 {
		HandleError(c, http.StatusInternalServerError, errors.New(result.Errmsg))
		return
	}
	wxUserAdd :=structure.WxUserAdd{
		Openid: result.Openid,
		SessionKey: result.SessionKey,
	}
	err = model.WxLogin(wxUserAdd)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, result)
}