package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:xianye@tcp(localhost)/db1")
	if err != nil {
		panic(err)
	}

	dB = db
} //准备数据库抽象
