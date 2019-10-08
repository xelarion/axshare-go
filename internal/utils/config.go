package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
