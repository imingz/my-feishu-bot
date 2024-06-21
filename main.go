package main

import (
	"log/slog"
	"os"
	"xiaoxiaojiqiren/internal/bot"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/slogor"
)

func main() {
	// 初始化机器人
	b := bot.NewBot()

	if env.Active == env.PRO {
		b = bot.NewBot(bot.WithLogger(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
			TimeFormat: "2006-01-02 15:04:05.000000",
			ShowSource: true,
			Level:      slog.LevelInfo,
			NoColor:    true,
		}))))
	}

	// 运行机器人
	b.Run()
}
