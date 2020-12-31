package utils

import (
	"goScaffold/configs"
	"goScaffold/database"
	"time"

	"github.com/iris-contrib/middleware/jwt"
)

type jwtToken struct {
	Secret        string
	JwtMiddleware *jwt.Middleware
}

func (imp *jwtToken) init() {
	JWT.Secret = configs.ConfigUtil.LoginSecret
	imp.JwtMiddleware = jwt.New(jwt.Config{
		// ErrorHandler: func(ctx context.Context, err error) {
		// 	if err == nil {
		// 		return
		// 	}
		// 	ctx.StopExecution()
		// 	ctx.StatusCode(iris.StatusUnauthorized)
		// 	ctx.JSON(model.ErrorUnauthorized(err))
		// },

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(imp.Secret), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}

func (imp *jwtToken) GenerateJwtToken(user *database.User) string {
	jwtToken := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iss": "Tango",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(2 * time.Hour * time.Duration(1)).Unix(),
	})
	token, _ := jwtToken.SignedString([]byte(imp.Secret))
	return token
}

var JWT jwtToken
