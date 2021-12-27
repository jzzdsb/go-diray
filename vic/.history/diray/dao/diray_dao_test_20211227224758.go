/*
 * @Author: vic
 * @Date: 2021-12-27 22:19:26
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 22:47:54
 * @Description:
 */
package dao

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("query", func(t *testing.T) {
		diary := Query("select d.id, d.type as type_info, d.create_date, d.create_user,  dt.txt from diary  d left join diary_txt dt on d.txt_id = dt.id")
		fmt.Printf("%v", diary)
	})
}

func TestInsert(t *testing.T) {
	t.Run("insert", func(t *testing.T) {
		Insert("insert into diary_txt(text)vaule(?)")
		Insert("insert into diary(type,txt_id)value(0,1)")
	})
}
