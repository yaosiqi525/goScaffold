package main

import (
	"goScaffold/router"

	_ "goScaffold/utils"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	// 初始化

	// 初始化日志
	log.LoadConfiguration("./logs/log.json")

	// 设置路由
	router.SetRouter(app)

	app.Use(recover.New())

	app.Run(iris.Addr(":80"), iris.WithoutServerError(iris.ErrServerClosed))
}
