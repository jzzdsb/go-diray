/*
 * @Author: vic
 * @Date: 2021-12-27 22:19:26
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 23:25:34
 * @Description:
 */
package test

import (
	"fmt"
	"testing"
	"org/vic"
)

func TestQuery(t *testing.T) {
	t.Run("query", func(t *testing.T) {
		diary := Query("select d.id, d.type as type_info, d.create_date, d.create_user,  dt.txt from diary  d left join diary_txt dt on d.txt_id = dt.id")
		fmt.Printf("%v", diary)
	})
}

func TestInsert(t *testing.T) {
	t.Run("insert", func(t *testing.T) {
		uuid := utils.GetUuid()
		Insert("insert into diary_txt(id,text)value(?,?)", uuid, "今天把dao层写完我了，争取明天可以通过前端来写日记，现阶段心情依然不好")
		//Insert("insert into diary(type,txt_id)value(?,?)",0,1)
	})
}
