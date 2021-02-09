package configs

import (
	"fmt"
	"os"

	log "github.com/jeanphorn/log4go"
	"gopkg.in/yaml.v2"
)

var ConfigUtil ConfigStruct

func init() {
	if ConfigUtil.init == 0 {
		Config, err := readEnvConfig()
		if err != nil {
			return
		}
		Config.init = 1
		ConfigUtil = Config
		log.Info("config: %+v", ConfigUtil)
	}
	return
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type ConfigStruct struct {
	init        int
	DBConfig    DBConfig `yaml:"db"`
	LoginSecret string   `yaml:"secret"`
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
