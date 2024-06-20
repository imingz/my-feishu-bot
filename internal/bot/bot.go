package bot

import (
	"context"
	"log/slog"
	"xiaoxiaojiqiren/config"
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/bot/ws"
	huihutongclient "xiaoxiaojiqiren/internal/pkg/HuiHutong_client"
)

type Bot struct {
	config     *config.Config // 配置
	larkClient *lark.Client   // Lark 客户端
	wsClient   *ws.Client     // 事件订阅客户端
}

func NewBot() *Bot {
	// 初始化配置
	config := config.NewConfig()
	slog.Debug("初始化配置", "config", config)

	// 初始化慧湖通客户端
	huihutong := huihutongclient.NewClient()

	// 初始化 Lark 客户端
	larkClient := lark.NewClient(config.Bot.AppID, config.Bot.AppSecret, huihutong)

	// 初始化事件订阅客户端
	wsClient := ws.NewClient(config.Bot.AppID, config.Bot.AppSecret, larkClient)

	return &Bot{
		config:     config,
		wsClient:   wsClient,
		larkClient: larkClient,
	}
}

func (b *Bot) Run() {
	err := b.wsClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
