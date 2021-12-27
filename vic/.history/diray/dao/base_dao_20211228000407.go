/*
 * @Author: vic
 * @Date: 2021-12-28 00:03:35
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:03:36
 * @Description:
 */
package dao

import(
	
)


/**
 * @description: 基础查询数据
 * @param {*} string 0 执行sql 后参数
 * @return {*}
 */
 func Query(sql string, values ...interface{}) []model.Diary {
	var diarys []model.Diary
	row := dataconfig.DB.Raw(sql, values...)
	if row.Error != nil {
		panic("查询sql有误")
	}
	row.Scan(&diarys)
	return diarys
}

/**
 * @description: 基础插入数据
 * @param {string} sql
 * @param {...interface{}} values
 * @return {*}
 */
func Insert(sql string, values ...interface{}) {
	row := dataconfig.DB.Raw(sql, values...).Row()
	if row.Err() != nil {
		panic("数据插入失败")
	}
}
