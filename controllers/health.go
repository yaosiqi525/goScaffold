package controllers

import (
	"github.com/kataras/iris/v12"
)

type health struct{}

func (imp *health) Index(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(ApiResource(SUCCESS, ctx, "ok"))
	return
}

var Health health
