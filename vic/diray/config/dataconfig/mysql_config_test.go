/*
 * @Author: vic
 * @Date: 2021-12-27 01:43:02
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 02:28:51
 * @Description:
 */
package dataconfig

import (
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		db := InitDB()
		diary := Query(db, "select d.id, d.type as type_info, d.create_date, d.create_user,  dt.txt from diary  d left join diary_txt dt on d.txt_id = dt.id")
		t.Logf("diary:%v", diary)
		db.Close()
	})
}
