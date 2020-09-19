package main

import (
	"goScaffold/utils"
	"goScaffold/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	// 初始化
	utils.Init()

	// 设置路由
	router.SetRouter(app)

	app.Run(iris.Addr(":80"))
}
