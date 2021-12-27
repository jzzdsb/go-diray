/*
 * @Author: vic
 * @Date: 2021-12-27 17:50:39
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:39:03
 * @Description: 通过gorm链接数据库
 */
package dataconfig

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//数据库信息
const (
	NAME     = "root"
	PASSWORD = "835512239"
	URL      = "127.0.0.1:3306"
	DATABASE = "public"
)

/**
 * @description: 初始化数据库连接
 */
func init() {
	dsn := NAME + ":" + PASSWORD + "@tcp(" + URL + ")/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	//关闭事务
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	if err != nil {
		panic("数据库连接失败")
	}
}
