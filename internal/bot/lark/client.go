package lark

import (
	"xiaoxiaojiqiren/config"
	huihutongclient "xiaoxiaojiqiren/internal/pkg/HuiHutong_client"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/mail"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type Client struct {
	config    *config.Config
	client    *lark.Client
	Huihutong *huihutongclient.Client
	Mail      *mail.Client
}

func NewClient(config *config.Config) *Client {
	logLevel := larkcore.LogLevelInfo
	if env.Active == env.DEV {
		logLevel = larkcore.LogLevelDebug
	}

	// 初始化慧湖通客户端
	huihutong := huihutongclient.NewClient()

	return &Client{
		client:    lark.NewClient(config.Bot.AppID, config.Bot.AppSecret, lark.WithLogLevel(logLevel)),
		Huihutong: huihutong,
		config:    config,
		Mail:      mail.New(&config.Mail),
	}
}
