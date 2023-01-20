package mapper

//@author by Hchier
//@Date 2023/1/20 20:24

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Connect("mysql", "root:pyh903903@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("数据库连接失败：", err.Error())
	}
	//最大空闲连接
	Db.SetMaxIdleConns(20)
	//最大连接
	Db.SetMaxOpenConns(100)
}
