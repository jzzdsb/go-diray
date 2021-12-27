/*
 * @Author: vic
 * @Date: 2021-12-27 17:50:39
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:08:07
 * @Description: 通过gorm链接数据库
 */
package dataconfig

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	NAME     = "root"
	PASSWORD = "835512239"
	URL      = "127.0.0.1:3306"
	DATABASE = "public"
)

/**
 * @description: 初始化数据库
 * @param {*}
 * @return {*}
 */
func init() {
	dsn := NAME+":"+PASSWORD+"@tcp("+URL+")/"+DATABASE+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),  &gorm.Config{
		SkipDefaultTransaction: true,
	  })
	DB = db
	if err != nil {
		panic("数据库连接失败")
	}
}
