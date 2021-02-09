package configs

import (
	"fmt"
	"os"

	log "github.com/jeanphorn/log4go"
	"gopkg.in/yaml.v2"
)

var ConfigUtil ConfigStruct

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type OSSConfigs struct {
	EndPoint        string `yaml:"endPoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	BaseUrl         string `yaml:"baseUrl"`
}

type ConfigStruct struct {
	init        int
	DBConfig    DBConfig   `yaml:"db"`
	OSSConfig   OSSConfigs `yaml:"oss"`
	LoginSecret string     `yaml:"secret"`
}

func getEnv() string {
	fmt.Println("env: ", os.Args[1])
	return os.Args[1]
}

func readEnvConfig() (ConfigStruct, error) {
	var configS ConfigStruct
	f, err := os.Open("./configs/" + getEnv() + ".yml")
	defer f.Close()

	fDecode := yaml.NewDecoder(f)
	fDecode.Decode(&configS)
	return configS, err
}

func InitConfig() (ConfigStruct, error) {
	if ConfigUtil.init == 0 {
		Config, err := readEnvConfig()
		if err != nil {
			return Config, err
		}
		Config.init = 1
		ConfigUtil = Config
		log.Info("config: %+v", ConfigUtil)
	}
	return ConfigUtil, nil
}
