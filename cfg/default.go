package cfg

import "github.com/spf13/viper"

func setDefaultConfig() {
	viper.SetDefault("LogConfig.EnableConsole", true)
	viper.SetDefault("LogConfig.ConsoleJSONFormat", false)
	viper.SetDefault("LogConfig.ConsoleLevel", "info")
	viper.SetDefault("LogConfig.EnableFile", true)
	viper.SetDefault("LogConfig.FileJSONFormat", true)
	viper.SetDefault("LogConfig.FileLevel", "debug")
	viper.SetDefault("LogConfig.FileLocation", "./gfmt.log")
	viper.SetDefault("DBConf.Host", "127.0.0.1")
	viper.SetDefault("DBConf.Port", "3306")
	viper.SetDefault("DBConf.User", "root")
	viper.SetDefault("DBConf.Password", "root")
	viper.SetDefault("DBConf.Name", "lums")
	viper.SetDefault("Env", "prod")
	viper.SetDefault("Port", "9876")
	viper.SetDefault("Verbose", false)
}
