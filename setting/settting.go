package setting

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Name      string       `json:"name"`
	Mode      string       `json:"mode"`
	Port      int          `json:"port"`
	JwtKey    string       `json:"jwt_key"`
	LogConfig *LogConfig   `json:"log_config"`
	SqlConfig *MySQLConfig `json:"mysql_config"`
	Admin     *AdminConf   `json:"admin"`
}

type LogConfig struct {
	Path string `json:"path"`
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int16  `json:"port"`
	User     string `json:"user"`
	DB       string `json:"db"`
	Password string `json:"password"`
}

type AdminConf struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

var (
	Config *AppConfig
)

// 从配置文件中载入json字符串
func LoadConfig(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config failed: ", err)
	}
	err = json.Unmarshal(buf, &Config)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	InitMysql(Config.SqlConfig)
}
