package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

func CheckUserInfo(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	id := int64(token["id"].(float64))
	log.Info("id: %d", id)

	ctx.Values().Set("user_id", id)
	ctx.Next() // 执行下一个处理器。
}
