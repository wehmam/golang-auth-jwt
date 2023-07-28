package env 

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

func init() {
	fileConfig := "config.json"
	viper.SetConfigName(fileConfig)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while reading config file %s" , err))
	}
}

func String(key string, defaultValue string) string {
	value, exist := viper.Get(key).(string)
	if !exist {
		log.Fatal(fmt.Sprintf("Errow while find config file %s", "Invalid type assertion - from = " +key))
		return defaultValue
	}

	return value
}