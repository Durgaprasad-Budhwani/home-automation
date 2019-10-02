package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper(filePath string) *viper.Viper {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filePath)
	// read in config file and check your errors
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	// confirm where the file has been read in from
	fmt.Println(viper.ConfigFileUsed())
	return viper
}
