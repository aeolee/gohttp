package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type dbObj struct {
	db *sql.DB
}
type DbWorker struct {
	Dsn string
}
type countInfo struct {
	ip          string
	mac         string
	channelName string
	starTime    string
	endTime     string
	enter       int32
	leave       int32 //xml中使用exit表示离开人数，这里改用leave表示。
}

func init() {
	//
}

//简单的处理下错误，便于调试
func checkErr(err error, str string) {
	if err != nil {
		fmt.Printf("%s,err:%v", str, err)
		panic(err)
	}
}

func (d *dbObj) Open() *sql.DB {
	var err error
	dbw := DbWorker{
		Dsn: "aeo:mp3abc@@@tcp(cd-cdb-99uyhax8.sql.tencentcdb.com:63118)/counting?charset=utf8"}
	d.db, err = sql.Open("mysql", dbw.Dsn)

	checkErr(err, "连接数据库失败")

	return d.db
}

func Select() {
	dbc := &dbObj{}
	db := dbc.Open()
	defer db.Close()

	stmt, _ := db.Prepare("select ip, mac,channelName,starTime,endTime,enter,`leave` " +
		"from people where ip>?")
	rows, _ := stmt.Query("192.168.1.1")
	defer rows.Close()

	count := &countInfo{}
	for rows.Next() {
		err := rows.Scan(&count.ip, &count.mac, &count.channelName,
			&count.starTime, &count.endTime, &count.enter, &count.leave)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%s 到 %s  进入人数:%4d    离开人数：%4d",
			count.starTime, count.endTime, count.enter, count.leave)
	}
}

func Insert(info countInfo) {
	dbc := &dbObj{}
	db := dbc.Open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO people SET ip=?,mac=?," +
		"channelName=?,starTime=?,endTime=?,enter=?,`leave`=?")
	checkErr(err, "sql语句有语法错误")
	result, err := stmt.Exec(info.ip, info.mac, info.channelName, info.
		starTime, info.endTime, info.enter, info.leave)
	checkErr(err, "插入数据失败")
	rowsaffected, err := result.RowsAffected()
	checkErr(err, "获取受影响行数失败")
	fmt.Println("受影响行数：", rowsaffected)
}

/*func main() {
	info := countInfo{ip:"192.168.1.31",mac:"10:12:fb:de:0b:00", channelName:"测试门",
		starTime:"2021-06-25 15:10:00",endTime:"2021-06-25 15:15:00",enter:123,exit:101}
	Insert(info)

	Select()
}
*/
