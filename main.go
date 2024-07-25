package main

import (
	"cell_phone_store/model"
	"cell_phone_store/structure"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"cell_phone_store/router"
	"log"
)

func main()  {
	//读取配置文件
	var conf structure.Config
	if _, err := toml.DecodeFile(structure.ConfigPath, &conf); err != nil {
		log.Fatalf("Error db config file: %v", err)
	}
	//初始化数据库
	if err := model.InitMysql(conf.Database); err != nil {
		log.Print("数据库连接失败", err)
	}
	//初始化服务
	ginServer :=gin.Default()
	//注册路由
	router.Router(ginServer)
	//服务端口
	ginServer.Run(":8666")
} 