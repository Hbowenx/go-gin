package controller

import (
	"cell_phone_store/model"
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//资讯列表
func InformationList(c *gin.Context) {
	var informationListRequest request.InformationListRequest
	if err := c.ShouldBindWith(&informationListRequest,binding.Form); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	informationList,err :=model.InformationList(informationListRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, informationList)
}

//新增资讯
func InformationAdd(c *gin.Context) {
	var informationAddRequest request.InformationAddRequest
	if err := c.ShouldBindWith(&informationAddRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	information := structure.InformationAdd{
		MasterMap: informationAddRequest.MasterMap,
		Title:     informationAddRequest.Title,
		Text:      informationAddRequest.Text,
		Images:    informationAddRequest.Images,
	}
	if err := model.InformationAdd(information); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK,"")
}

//资源详情
func InformationInfo(c *gin.Context)  {
	var InformationInfoRequest request.InformationInfoRequest
	if err := c.ShouldBindWith(&InformationInfoRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	informationInfo,err :=model.InformationInfo(InformationInfoRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK,informationInfo)
}

//资源删除
func InformationDel(c *gin.Context)  {
	var InformationDelRequest request.InformationDelRequest
	if err := c.ShouldBindWith(&InformationDelRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	err :=model.InformationDel(InformationDelRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK,"")
}