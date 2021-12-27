/*
 * @Author: vic
 * @Date: 2021-12-27 22:04:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-28 00:26:58
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
func WriteDiary(mood int, txt string) error {
	txt_id := utils.GetUuid()
	diary_id := utils.GetUuid()
	tx := dataconfig.DB.Begin()

	defer func() {
		if r:=crecover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Raw("insert into diary_txt(id,txt)value(?,?)", txt_id, txt).Row().Err(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Raw("insert into diary(id,type,txt_id)value(?,?,?)", diary_id, mood, txt_id).Row().Err(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
