package bot

import (
	"context"
	"log/slog"
	"xiaoxiaojiqiren/config"
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/bot/ws"
)

type Bot struct {
	Config     *config.Config // 配置
	LarkClient *lark.Client   // Lark 客户端
	WsClient   *ws.Client     // 事件订阅客户端
}

func NewBot() *Bot {
	// 初始化配置
	config := config.NewConfig()
	slog.Debug("初始化配置", "config", config)

	// 初始化 Lark 客户端
	larkClient := lark.NewClient(config.Bot.AppID, config.Bot.AppSecret)

	// 初始化事件订阅客户端
	wsClient := ws.NewClient(config.Bot.AppID, config.Bot.AppSecret, larkClient)

	return &Bot{
		Config:     config,
		WsClient:   wsClient,
		LarkClient: larkClient,
	}
}

func (b *Bot) Run() {
	err := b.WsClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
