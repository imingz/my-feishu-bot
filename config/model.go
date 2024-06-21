package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Bot         BotConfig         `yaml:"Bot"`
	ConfigTable ConfigTableConfig `yaml:"ConfigTable"`
	Mail        MailConfig        `yaml:"Mail"`
}

// BotConfig 包含 Bot 的详细配置
type BotConfig struct {
	AppID             string `yaml:"AppID"`
	AppSecret         string `yaml:"AppSecret"`
	VerificationToken string `yaml:"VerificationToken"`
	EventEncryptKey   string `yaml:"EventEncryptKey"`
}

// ConfigTableConfig
type ConfigTableConfig struct {
	AppToken string `yaml:"AppToken"`
	TableId  string `yaml:"TableId"`
	ViewId   string `yaml:"ViewId"`
}

// MailConfig
type MailConfig struct {
	SMTP struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"SMTP"`
}

// 实现 Config 的 String() 方法
func (c Config) String() string {
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("BotConfig: error marshalling to JSON: %v", err)
	}
	return string(data)
}
