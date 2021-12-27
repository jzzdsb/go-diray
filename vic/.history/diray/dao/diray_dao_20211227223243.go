/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 22:32:09
 * @Description: 日记相关数据库操作
 */
package dao

import (
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
func insert(sql string,values ...interface{})