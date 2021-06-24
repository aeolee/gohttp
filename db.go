package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "src/github.com/go-sql-driver/mysql"
)

type dbObj struct {
	db *sql.DB
}
type userInfo struct {
	uid        int32
	username   string
	department string
	created    string
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
	d.db, err = sql.Open("mysql", "godbdemo:mp3abc@tcp(localhost:3306)/godbdemo?charset=utf8")
	checkErr(err, "连接数据库失败")

	return d.db
}

func Select() {
	dbc := &dbObj{}
	db := dbc.Open()
	defer db.Close()

	stmt, _ := db.Prepare("select department,`username` from userinfo where uid>?")
	rows, _ := stmt.Query(99)
	defer rows.Close()

	user := &userInfo{}

	for rows.Next() {
		err := rows.Scan(&user.username, &user.department)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user.username, ":", user.department)
	}
}

func Insert(info userInfo) {
	dbc := &dbObj{}
	db := dbc.Open()
	defer db.Close()

	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	result, err := stmt.Exec(info.username, info.department, info.created)
	checkErr(err, "插入数据失败")
	rowsaffected, err := result.RowsAffected()
	checkErr(err, "获取受影响行数失败")
	fmt.Println("受影响行数：", rowsaffected)
}

func main() {

	//info := userInfo{uid:0,username:"stone",department:"HR",created: "20210625"}
	//Insert(info)

	Select()

	/*stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	for i := 1; i < 100; i++ {
		res, err := stmt.Exec("技术", "研发部门", "20210624")
		checkErr(err)
		fmt.Println("技术 i==", i, "res===", res.RowsAffected())
	}

	res, err := stmt.Exec("技术","研发部门","20210624")
	checkErr(err)

	tests,err := res.RowsAffected()
	checkErr(err)

	fmt.Println(tests)

	db.Close()*/
}
