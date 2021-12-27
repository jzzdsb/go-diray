/*
 * @Author: vic
 * @Date: 2021-12-27 21:54:10
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 23:13:22
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
	Id         string
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


type DiaryText struct {
	Id  string
	Txt string
}

func (DiaryText) TableName() string {
	return "diary_txt"
}
