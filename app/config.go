package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	if gin.Mode() == gin.ReleaseMode {
		Config.SetConfigName("config")
	} else {
		Config.SetConfigName("config-" + gin.Mode())
	}
	Config.SetConfigType("json")
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
