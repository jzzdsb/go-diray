/*
 * @Author: vic
 * @Date: 2021-12-27 22:19:26
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 22:29:14
 * @Description:
 */
package dao

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		diary := Query("select d.id, d.type as type_info, d.create_date, d.create_user,  dt.txt from diary  d left join diary_txt dt on d.txt_id = dt.id", "")
		fmt.Printf("%v", diary)
	})
}
