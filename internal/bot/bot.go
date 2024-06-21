package bot

import (
	"context"
	"log/slog"
	"os"
	"xiaoxiaojiqiren/config"
	cron "xiaoxiaojiqiren/internal/bot/cron_client"
	ginclient "xiaoxiaojiqiren/internal/bot/gin_client"
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/bot/ws"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/slogor"
)

type Bot struct {
	logger    *slog.Logger      // 日志
	wsClient  *ws.Client        // 事件订阅客户端
	ginclient *ginclient.Client // gin 客户端
	cron      *cron.Client
}

var defaultLogger = slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
	TimeFormat: "2006-01-02 15:04:05.000000",
	ShowSource: true,
	Level:      slog.LevelDebug,
	NoColor:    env.Active == env.PRO,
}))

func NewBot(opts ...Option) *Bot {
	// 初始化配置
	config := config.NewConfig()
	slog.Info("初始化配置", "config", config)

	// 初始化 Lark 客户端
	larkClient := lark.NewClient(config)

	b := &Bot{
		logger:    defaultLogger,
		wsClient:  ws.NewClient(config.Bot.AppID, config.Bot.AppSecret, larkClient),
		ginclient: ginclient.NewClient(config.Bot.VerificationToken, config.Bot.EventEncryptKey, larkClient),
		cron:      cron.New(),
	}

	for _, opt := range opts {
		opt(b)
	}

	slog.SetDefault(b.logger)
	return b
}

func (b *Bot) Run() {
	// 启动 gin 客户端
	go func() {
		b.ginclient.Run()
	}()

	// 启动 cron 定时任务
	go func() {
		b.cron.Start()
	}()

	// 启动事件订阅客户端
	err := b.wsClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
