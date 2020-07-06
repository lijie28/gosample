package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"app/tools"
	"app/userdb"
)

func main() {
	fmt.Println("go my sql")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())

	}

	defer db.Close()

	tools.SayHallo()
	userdb.Init()

}
