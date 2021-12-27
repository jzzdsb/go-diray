/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:22:56
 * @Description: 日记相关数据库操作
 */
package dao

import (
	"org/vic/diray/config/dataconfig"
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
	//开启事务
	dataconfig.DB.Begin()
	Insert("insert into diary_txt(id,txt)value(?,?)", txt_id, txt)
	Insert("insert into diary(id,type,txt_id)value(?,?,?)", diary_id, mood, txt_id)
}

func WriteDiary(mood int, txt string) error {
	txt_id := utils.GetUuid()
	diary_id := utils.GetUuid()
	tx := dataconfig.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Raw("insert into diary_txt(id,txt)value(?,?)", txt_id, txt).Row(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
