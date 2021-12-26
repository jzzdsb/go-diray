/*
 * @Author: vic
 * @Date: 2021-12-27 00:48:02
 * @LastEditors: vic
 * @LastEditTime: 2021-12-27 02:30:59
 * @Description:mysql
 */

package dataconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Diary struct {
	id          int64
	type_info   int8
	create_date string
	create_user string
	txt         string
}

func InitDB() *sql.DB {
	DB, _ := sql.Open("mysql", "root:835512239@tcp(localhost:3306)/public?charset=utf8")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}

	fmt.Println("connnect success")
	return DB
}

//查询操作
func Query(db *sql.DB, query string) Diary {
	var diaryInfo Diary
	rows, e := db.Query(query)
	checkErr(e)

	for rows.Next() {
		e := rows.Scan(&diaryInfo.id, &diaryInfo.type_info, &diaryInfo.create_date, &diaryInfo.create_user, &diaryInfo.txt)
		checkErr(e)
	}
	rows.Close()
	return diaryInfo
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
