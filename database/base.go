package database

import (
	"fmt"
	"goScaffold/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"xorm.io/core"
)

// DBUtils 数据库单例
var DBUtils *xorm.Engine

// Init 数据库初始化
func Init(config configs.DBConfig) (err error) {
	// 数据库初始化
	var dbString string
	// root:123@(127.0.0.1:3306)/test?charset=utf8
	dbString = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", config.User, config.Password, config.Host, config.Port, config.Database, config.Charset)

	DBUtils, err = xorm.NewEngine("mysql", dbString)
	// 打开Sql日志
	DBUtils.ShowSQL(true)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	DBUtils.SetTableMapper(tbMapper)
	initDbTable()
	return
}

func initDbTable() {
	DBUtils.Sync2(new(User))
}
