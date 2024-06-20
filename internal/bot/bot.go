package bot

import (
	"context"
	"log/slog"
	"xiaoxiaojiqiren/config"
	ginclient "xiaoxiaojiqiren/internal/bot/gin_client"
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/bot/ws"
	huihutongclient "xiaoxiaojiqiren/internal/pkg/HuiHutong_client"
)

type Bot struct {
	wsClient  *ws.Client        // 事件订阅客户端
	ginclient *ginclient.Client // gin 客户端
}

func NewBot() *Bot {
	// 初始化配置
	config := config.NewConfig()
	slog.Info("初始化配置", "config", config)

	// 初始化慧湖通客户端
	huihutong := huihutongclient.NewClient()

	// 初始化 Lark 客户端
	larkClient := lark.NewClient(config.Bot.AppID, config.Bot.AppSecret, huihutong)

	// 初始化事件订阅客户端
	wsClient := ws.NewClient(config.Bot.AppID, config.Bot.AppSecret, larkClient)

	// 初始化 gin 客户端
	ginclient := ginclient.NewClient(config.Bot.VerificationToken, config.Bot.EventEncryptKey, larkClient)

	return &Bot{
		wsClient:  wsClient,
		ginclient: ginclient,
	}
}

func (b *Bot) Run() {
	// 启动 gin 客户端
	go func() {
		b.ginclient.Run()
	}()

	// 启动事件订阅客户端
	err := b.wsClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
