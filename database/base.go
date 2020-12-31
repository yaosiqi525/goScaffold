package database

import (
	"fmt"
	"goScaffold/configs"
	"log"

	_ "github.com/go-sql-driver/mysql"
	logger "github.com/jeanphorn/log4go"
	"github.com/xormplus/xorm"
	"xorm.io/core"
)

// DBUtils 数据库单例
var DBUtils *xorm.Engine

// Init 数据库初始化
func init() {
	config := configs.ConfigUtil.DBConfig
	// 数据库初始化
	var dbString string
	// root:123@(127.0.0.1:3306)/test?charset=utf8
	dbString = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", config.User, config.Password, config.Host, config.Port, config.Database, config.Charset)

	var err error
	DBUtils, err = xorm.NewEngine("mysql", dbString)
	if err != nil {
		log.Fatal(err)
	}
	// 打开Sql日志
	DBUtils.ShowSQL(true)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	DBUtils.SetTableMapper(tbMapper)
	initDbTable()
	logger.Info("init db success")
	return
}

func initDbTable() {
	err := DBUtils.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}
}
