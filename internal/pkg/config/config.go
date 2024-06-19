package config

import (
	"xiaoxiaojiqiren/internal/pkg/env"

	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	APP    AppConfig    `yaml:"App"`
	Qrcode QrcodeConfig `yaml:"qrcode"`
}

type AppConfig struct {
	ID                string `yaml:"ID"`
	Secret            string `yaml:"Secret"`
	VerificationToken string `yaml:"VerificationToken"`
	EventEncryptKey   string `yaml:"EventEncryptKey"`
}

type QrcodeConfig struct {
	OpenId   string `yaml:"openId"`
	BaseUrl  string `yaml:"baseUrl"`
	ChatId   string `yaml:"chatId"`
	ThreadId string `yaml:"threadId"`
}

func init() {
	viper.SetConfigName(env.Active) // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")        // 还可以在工作目录中查找配置
	viper.AddConfigPath("./configs")
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
}

func Get() Config {
	return *config
}
