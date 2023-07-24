package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error Reading the configs")
	}

	// After Config file found and successfully parsed

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
