package common

//@author by Hchier
//@Date 2023/1/26 21:58
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	// 这个init的优先级比较高，又必须在连接mysql前加载完配置，于是只能放这里了
	LoadConfig("config.conf")

	var err error
	Db, err = sqlx.Connect(DriverName, DataSourceName)
	if err != nil {
		ErrLog("数据库连接失败：", err.Error())
	}
	//最大空闲连接
	Db.SetMaxIdleConns(20)
	//最大连接
	Db.SetMaxOpenConns(100)
}
