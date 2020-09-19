package utils

import (
	"goScaffold/configs"
	"goScaffold/database"

	log "github.com/jeanphorn/log4go"
)

func Init() error {
	// 初始化日志
	log.LoadConfiguration("./logs/log.json")

	// 初始化配置文件
	config, err := configs.InitConfig()
	if err != nil {
		log.Error("init config error:", err.Error())
	}

	// 数据库初始化
	err = database.Init(config.DBConfig)
	if err != nil {
		return err
	}
	log.Info("init db success")

	// 初始化JWT
	JWT.Secret = config.LoginSecret
	JWT.InitJWT()

	log.Info("config: %+v", config)

	log.Info("init success")
	return nil
}
