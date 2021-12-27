/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 22:57:52
 * @Description: 日记相关数据库操作
 */
package dao

import (
	"fmt"
	"org/vic/diray/config/dataconfig"
	"org/vic/diray/model"
)

/**
 * @description: 查询数据
 * @param {*} string 0 执行sql 后参数
 * @return {*}
 */
func Query(sql string, values ...interface{}) []model.Diary {
	var diarys []model.Diary
	dataconfig.DB.Raw(sql, values...).Scan(&diarys)
	return diarys
}

/**
 * @description: 插入数据
 * @param {string} sql
 * @param {...interface{}} values
 * @return {*}
 */
func Insert(sql string, values ...interface{}) {
	row:= dataconfig.DB.Raw(sql, values...).Row()
	
}
