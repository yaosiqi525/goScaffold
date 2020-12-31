package router

import (
	"goScaffold/controllers"
	"goScaffold/middlewares"
	"goScaffold/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func cors(ctx iris.Context) {
	middlewares.Cors(ctx)
}

// SetRouter 设置路由
func SetRouter(app *iris.Application) {
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(controllers.ApiResource(controllers.PAGE_NOT_FOUND, nil, "404 Not Found"))
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})

	app.Use(cors)

	app.Get("/health", controllers.Health.Index)

	app.PartyFunc("/api", func(mainRoute iris.Party) {
		// 无需权限
		mainRoute.Post("/register", controllers.Login.Register)
		mainRoute.Post("/login", controllers.Login.Login)

		// 需要登陆权限
		mainRoute.Get("/userinfo", utils.JWT.JwtMiddleware.Serve, middlewares.CheckUserInfo, controllers.Login.Index)
	})

	app.Options("*", func(ctx iris.Context) {
		ctx.StatusCode(200)
	})

}
