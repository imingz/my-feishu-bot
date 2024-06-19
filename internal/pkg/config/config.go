package config

import (
	"path"
	"runtime"
	"xiaoxiaojiqiren/internal/pkg/env"

	"github.com/spf13/viper"
)

var config = new(Config)

func NewConfig() *Config {
	viper.SetConfigName(env.Active)                        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")                            // 如果配置文件的名称中没有扩展名，则需要配置此项
	_, filename, _, _ := runtime.Caller(0)                 // 获取当前文件（config.go）路径
	confPath := path.Join(path.Dir(filename), "./configs") // 获取当前文件目录
	viper.AddConfigPath(confPath)                          // 添加配置文件路径
	viper.AddConfigPath(".")                               // 还可以在工作目录中查找配置
	viper.AddConfigPath("./configs")                       // 还可以在工作目录的 configs 目录中查找配置
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			panic("配置文件未找到" + err.Error())
		} else {
			// 配置文件被找到，但产生了另外的错误
			panic("配置文件解析错误" + err.Error())
		}
	}
	if err := viper.Unmarshal(config); err != nil {
		panic("配置文件解析错误" + err.Error())
	}
	return config
}

func Get() *Config {
	return config
}
