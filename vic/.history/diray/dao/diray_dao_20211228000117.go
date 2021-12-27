/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:00:25
 * @Description: 日记相关数据库操作
 */
package dao

import (
	"org/vic/diray/config/dataconfig"
	"org/vic/diray/model"
	"org/vic/diray/utils"
)

/**
 * @description: 查询数据
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
 * @description: 插入数据
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

/**
 * @description:写日记
 * @param {string} txt
 * @return {*}
 */
func WriteDiary(,txt string) {
	txt_id := utils.GetUuid()
	diary_id := utils.GetUuid()
	Insert("insert into diary_txt(id,txt)value(?,?)", txt_id, "今天把dao层写完我了，争取明天可以通过前端来写日记，现阶段心情依然不好")
	Insert("insert into diary(id,type,txt_id)value(?,?,?)", diary_id, 0, txt_id)
}
