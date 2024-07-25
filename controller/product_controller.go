package controller

import (
	"cell_phone_store/model"
	"cell_phone_store/request"
	"cell_phone_store/structure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//产品列表
func ProductList(c *gin.Context)  {
	var productListRequest request.ProductListRequest
	if err := c.ShouldBindWith(&productListRequest,binding.Form); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	productList,err :=model.ProductList(productListRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, productList)
}

//产品新增
func ProductAdd(c *gin.Context)   {
	var productAddRequest request.ProductAddRequest
	if err := c.ShouldBindWith(&productAddRequest,binding.Form); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	productAdd := structure.ProductAdd{
		Desc: productAddRequest.Desc,
		Images: productAddRequest.Images,
		Sort: productAddRequest.Sort,
		LabelId: productAddRequest.LabelId}
	if err := model.ProductAdd(productAdd); err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK, "")
}

//产品详情
func ProductInfo(c *gin.Context)   {
	var productInfoRequest request.ProductInfoRequest
	if err := c.ShouldBindWith(&productInfoRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	productInfo,err :=model.ProductInfo(productInfoRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK,productInfo)
}


//产品删除
func ProductDel(c *gin.Context)   {
	var ProductDelRequest request.ProductDelRequest
	if err := c.ShouldBindWith(&ProductDelRequest, binding.Form); err != nil {
		HandleError(c, http.StatusBadRequest, err)
		return
	}
	err :=model.ProductDel(ProductDelRequest)
	if err!=nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(c, http.StatusOK,"")
}