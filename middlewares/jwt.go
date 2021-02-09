package middlewares

import (
	"goScaffold/utils"

	"github.com/iris-contrib/middleware/jwt"
	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

// JwtCheck 解析jwt token
func JwtCheck(ctx iris.Context) {
	// imp.JwtMiddleware.
	err := utils.JWT.JwtMiddleware.CheckJWT(ctx)
	token := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	log.Info("token: %+v, err: %+v", token["id"].(float64), err)
	ctx.Next()
}
