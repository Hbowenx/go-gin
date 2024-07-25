package controller

import (
	"cell_phone_store/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取基础配置
func GetBaseConfig(c *gin.Context) {
	baseConfig,err :=model.GetBaseConfig()
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, baseConfig)
}

