package main

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

var viperConf *viper.Viper

func init() {
	viper.SetDefault("GithubAuthToken", "PLACEHOLDER")

	viper.AddConfigPath(path.Join(DIRS.ConfigHome, "script_manager"))
	viper.AddConfigPath(path.Join(DIRS.DataHome, "script_manager"))
	viper.SetConfigName("sm")

	err := viper.ReadInConfig()
	if err != nil {
		p := path.Join(DIRS.DataHome, "script_manager", "sm.toml");
		file, err := os.Create(p);
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		if err := file.Close(); err != nil {
			println(err.Error())
			os.Exit(1)
		}

		err = viper.ReadInConfig()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	viperConf = viper.GetViper()
}
