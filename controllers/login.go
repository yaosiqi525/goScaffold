package controllers

import (
	"goScaffold/database"
	"goScaffold/services"
	"goScaffold/utils"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

type login struct{}

func (imp *login) Index(ctx iris.Context) {
	id := ctx.Values().Get("user_id").(int64)
	user := services.User.GetUserInfo(id)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(ApiResource(SUCCESS, user, "ok"))
	return
}

type registerParams struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Register 注册
func (imp *login) Register(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	var params registerParams
	if err := ctx.ReadJSON(&params); err != nil {
		log.Info("error: %s", err.Error())
		ctx.JSON(ApiResource(PARAM_ERROR, nil, "参数异常"))
		return
	}
	// 检查是否已注册
	var registed database.User
	database.DBUtils.Where("phone = ?", params.Phone).Get(&registed)
	if registed.Name != "" {
		ctx.JSON(ApiResource(ERROR, nil, "已注册"))
		return
	}
	// 写入数据库
	registed.Name = params.Name
	registed.Phone = params.Phone
	registed.Password = params.Password
	database.DBUtils.Insert(registed)

	// 生成token
	token := utils.JWT.GenerateJwtToken(&registed)
	ctx.JSON(ApiResource(SUCCESS, token, "登陆成功"))
	return
}

// Login 注册
func (imp *login) Login(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	var params registerParams
	if err := ctx.ReadJSON(&params); err != nil {
		log.Info("error: %s", err.Error())
		ctx.JSON(ApiResource(PARAM_ERROR, nil, "参数异常"))
		return
	}
	// 检查是否已注册
	var registed database.User
	database.DBUtils.Where("phone = ? and password = ?", params.Phone, params.Password).Get(&registed)
	if registed.Name == "" {
		ctx.JSON(ApiResource(ERROR, nil, "未注册"))
		return
	}

	// 生成token
	token := utils.JWT.GenerateJwtToken(&registed)
	ctx.JSON(ApiResource(SUCCESS, token, "注册成功"))
	return
}

// Login 接入层单例
var Login login
