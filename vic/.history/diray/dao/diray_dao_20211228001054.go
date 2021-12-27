/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:10:45
 * @Description: 日记相关数据库操作
 */
package dao

import (
	"org/vic/diray/utils"
)

/**
 * @description:  写日记
 * @param {int} mood 心情
 * @param {string} txt 日记内容
 * @return {*}
 */
func WriteDiary(mood int, txt string) {
	txt_id := utils.GetUuid()
	diary_id := utils.GetUuid()
	//事务
	
	Insert("insert into diary_txt(id,txt)value(?,?)", txt_id, txt)
	Insert("insert into diary(id,type,txt_id)value(?,?,?)", diary_id, mood, txt_id)
}
