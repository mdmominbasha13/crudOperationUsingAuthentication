package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

var Appconfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("error while fetching path")
	}
	exepath := strings.Split(path, "\\")
	path = strings.Join(exepath[:len(exepath)-2], "\\")
	fmt.Println(exepath, path)
	config.AddConfigPath(path + "/pkg/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
