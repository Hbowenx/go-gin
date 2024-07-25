package router

import (
	"cell_phone_store/controller"
	"cell_phone_store/middleware"
	"github.com/gin-gonic/gin"
)

func Router(ginServer *gin.Engine)  {
	ginServer.GET("/getBaseConfig", controller.GetBaseConfig)

	//oss接口
	ossGroup := ginServer.Group("/oss")
	{
		ossGroup.Use(middleware.Authenticate).POST("/uploadOss", controller.UploadOss)
	}

	//微信接口
	wxGroup := ginServer.Group("/wx")
	{
		wxGroup.GET("/login", controller.WxLogin)
	}

	//用户接口
	userGroup := ginServer.Group("/user")
	{
		userGroup.POST("/login", controller.UserLogin)
		userGroup.GET("/info", controller.UserInfo)
	}

	//资讯接口
	// 资讯接口，不需要管理员权限
	informationGroup := ginServer.Group("/information")
	{
		informationGroup.GET("/list", controller.InformationList) // 资讯列表
		informationGroup.GET("/info", controller.InformationInfo) // 资讯详情
	}

	// 资讯接口，需要管理员权限
	adminInfoGroup := ginServer.Group("/information", middleware.AuthenticateAdmin)
	{
		adminInfoGroup.POST("/add", controller.InformationAdd) // 资讯新增
		adminInfoGroup.POST("/del", controller.InformationDel) // 资讯删除
	}

	//标签接口
	LabelGroup := ginServer.Group("/label")
	{
		LabelGroup.GET("/list", controller.LabelList)//标签列表
	}
	adminLabelGroup := ginServer.Group("/label",middleware.AuthenticateAdmin)
	{
		adminLabelGroup.POST("/add", controller.LabelAdd)//标签新增
		adminLabelGroup.POST("/del", controller.LabelDel)//标签删除
	}

	//产品接口
	ProductGroup := ginServer.Group("/product")
	{
		ProductGroup.GET("/list", controller.ProductList)//产品列表
		ProductGroup.GET("/info", controller.ProductInfo)//产品详情
	}
	adminProductGroup := ginServer.Group("/product",middleware.AuthenticateAdmin)
	{
		adminProductGroup.POST("/add", controller.ProductAdd)//产品新增
		adminProductGroup.POST("/del", controller.ProductDel)//产品删除
	}

}