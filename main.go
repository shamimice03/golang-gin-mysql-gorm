package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shamimice03/golang-gin-mysql-gorm/config"
	"github.com/spf13/viper"
)

func main() {

	configs := configSet()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func configSet() config.Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error Reading the configs")
	}

	// After Config file found and successfully parsed

	var configs config.Config

	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return configs
}
