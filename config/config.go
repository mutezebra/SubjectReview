package config

import (
	"os"

	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

var (
	Email  *email
	OSS    *oss
	Secret *secret
	Mysql  *mysql
	Redis  *redis
)

// InitConfig initializes the configuration for the project
// and unmarshall the configuration into the global variable "Conf"
func InitConfig() {
	wd, _ := os.Getwd()
	c := new(config)

	configDIR := os.Getenv("CONFIG_DIR")
	viper.AddConfigPath(configDIR) // auto

	viper.AddConfigPath(wd + "/config/")   // linux
	viper.AddConfigPath(wd + "\\config\\") // windows
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	configInit(c)
}

func configInit(c *config) {
	Email = &c.Email
	OSS = &c.OSS
	Secret = &c.Secret
	Mysql = &c.Mysql
	Redis = &c.Redis
}
