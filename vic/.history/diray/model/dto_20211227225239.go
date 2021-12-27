/*
 * @Author: vic
 * @Date: 2021-12-27 21:54:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 22:52:15
 * @Description:查询对象
 */
package model

import (
	"time"
)

/**
 * @description: 日记对象
 * @param {*}
 * @return {*}
 */
type Diary struct {
	Id         int64
	TypeInfo   int8
	CreateDate time.Time
	CreateUser string
	Txt        string
}

/**
 * @description: 设置表名
 * @param {*}
 * @return {*}
 */
func (Diary) TableName() string {
	return "diary"
}

type DiaryText struct{
	
}
