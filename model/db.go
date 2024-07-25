package model

import (
	"cell_phone_store/structure"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

//初始化数据库
func InitMysql(DatabaseConfig structure.DatabaseConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseConfig.User,
		DatabaseConfig.Password,
		DatabaseConfig.Host,
		DatabaseConfig.Port,
		DatabaseConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 尝试与数据库进行连接
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	// 设置最大空闲连接数、最大打开连接数等（可选）
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 保存 GORM 数据库实例到全局变量
	DB = db
	return nil
}