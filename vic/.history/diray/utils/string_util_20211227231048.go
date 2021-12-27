/*
 * @Author: vic
 * @Date: 2021-12-27 23:02:07
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 23:10:20
 * @Description:字符串相关方法
 */
package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GetUuid() string {
	return strings.ReplaceAll(uuid.New(),"-","");
}
