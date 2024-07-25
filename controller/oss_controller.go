package controller

import (
	"cell_phone_store/controller/oss"
	"cell_phone_store/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"path/filepath"
)

func UploadOss(c *gin.Context)  {
	var uploadOssRequest request.UploadOssRequest
	// 接收参数
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	if err := c.ShouldBindWith(&uploadOssRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	// 文件上传逻辑
	randName := RandString(8) // 假设 RandString 是你定义的生成随机字符串的函数
	fileType := filepath.Ext(header.Filename)
	fileName := uploadOssRequest.Dir + "/" + randName + fileType
	fileUrl,err := oss.UploadFile(file, fileName)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, fileUrl)
}
