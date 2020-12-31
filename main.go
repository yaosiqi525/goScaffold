package main

import (
	"goScaffold/router"

	_ "goScaffold/utils"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	// 初始化

	// 初始化日志
	log.LoadConfiguration("./logs/log.json")

	// 设置路由
	router.SetRouter(app)

	app.Run(iris.Addr(":80"))
}
