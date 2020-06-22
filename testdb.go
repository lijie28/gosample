package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbWorker struct {
	dsn      string
	db       *sql.DB
	userInfo usertb
}

type usertb struct {
	id int
	//NullString代表一个可为NULL的字符串。
	name sql.NullString
	//NullInt64代表一个可为NULL的Float64值。
	age sql.NullFloat64
}

func main() {
	var err error
	//初始化结构体，保存数据库参数
	dbw := dbWorker{
		dsn: "root:@/testsql?charset=utf8",
	}

	//打开数据库,并保存给结构体内db
	dbw.db, err = sql.Open("mysql", dbw.dsn)
	//如果打开失败，panic退出
	if err != nil {
		panic(err)
	}
	// log.Println("database successfully configured")
	// 关闭数据库
	defer dbw.db.Close()
	//插入数据
	dbw.insertData()
	//获取数据
	dbw.querData()
}

//创建方法，插入数据
func (dbw *dbWorker) insertData() {
	//预处理,插入数据
	//
	stmt, err := dbw.db.Prepare(`INSERT INTO tab_users(name,id,age) VALUES(?, null,?)`)

	if err != nil {
		fmt.Println(err)
		return
	}
	//函数执行完毕，关闭
	defer stmt.Close()

	//插入数据
	// var err error
	// var ret Result
	ret, err := stmt.Exec("hzxx", 29)
	fmt.Println("tesssst")
	if err != nil {
		fmt.Println("dberr", err)
		return
	}
	//获取id，执行成功，打印
	if LastInsertId, err := ret.LastInsertId(); err == nil {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	//获取行号，执行成功，打印
	if RowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

//初始化userInfo
func (dbw *dbWorker) querDataPre() {
	dbw.userInfo = usertb{}
}
func (dbw *dbWorker) querData() {

	//预处理,查询数据
	stmt, err := dbw.db.Prepare(`SELECT * From tab_users where age >= ? AND age < ?`)

	if err != nil {
		fmt.Println("querData stmt err:", err)
		return
	}
	defer stmt.Close()

	dbw.querDataPre()
	//取price20到30之间的数据
	rows, err := stmt.Query(20, 30)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		//取出数据库数据
		rows.Scan(&dbw.userInfo.id, &dbw.userInfo.name, &dbw.userInfo.age)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		//如果取出的数据为null,则赋一个0值
		if !dbw.userInfo.name.Valid {
			dbw.userInfo.name.String = ""
		}
		if !dbw.userInfo.age.Valid {
			dbw.userInfo.age.Float64 = 0.00
		}
		fmt.Println("get data,id:", dbw.userInfo.id, "name:", dbw.userInfo.name.String, "age", float64(dbw.userInfo.age.Float64))
	}
	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
