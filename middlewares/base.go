package middlewares

import (
	"bytes"
	"net/http"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

func ReadBody(resp *http.Request) string {
	resBody := resp.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resBody)
	return buf.String()
}

// Cors 实现服务端跨域
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, Content-Length,Accept-Encoding, X-CSRF-Token, AuthToken")
	if ctx.Request().Method == "OPTIONS" {
		ctx.StatusCode(200)
		return
	}
	// log.Info("request body: %s", ReadBody(ctx.Request()))
	ctx.StatusCode(iris.StatusOK)
	ctx.Next()
	log.Info("request end")
}
