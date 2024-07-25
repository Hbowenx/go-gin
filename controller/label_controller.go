package controller

import (
	"cell_phone_store/model"
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//标签列表
func LabelList(c *gin.Context)  {
	labelList,err :=model.LabelList()
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, labelList)
}

//标签新增
func LabelAdd(c *gin.Context)  {
	var labelAddRequest request.LabelAddRequest
	if err := c.ShouldBindWith(&labelAddRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	labelAdd := structure.LabelAdd{
		Name: labelAddRequest.Name,
		Sort: labelAddRequest.Sort,
	}
	err := model.LabelAdd(labelAdd)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, "")
}

//标签删除
func LabelDel(c *gin.Context)  {
	var labelDelRequest request.LabelDelRequest
	if err := c.ShouldBindWith(&labelDelRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	err := model.LabelDel(labelDelRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, "")
}