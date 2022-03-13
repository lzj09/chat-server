package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lzj09/chat-server/utils"
)

var MysqlClient *sqlx.DB

// Init 初始化mysql
func Init() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4,utf8&parseTime=true", utils.GetEnv("MYSQL_USERNAME", "root"), utils.GetEnv("MYSQL_PASSWORD", "123456"), utils.GetEnv("MYSQL_HOST", "127.0.0.1"), utils.GetEnv("MYSQL_PORT", "3306"), utils.GetEnv("MYSQL_DBNAME", "msg_db")))
	if err != nil {
		panic(err)
	}

	MysqlClient = db
}
