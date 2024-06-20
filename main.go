package main

import (
	"log/slog"
	"os"
	"xiaoxiaojiqiren/internal/bot"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/slogor"
)

func main() {
	// 设置 slog
	var slogLogLevel = slog.LevelInfo
	if env.Active == env.DEV {
		slogLogLevel = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: "2006-01-02 15:04:05.000",
		ShowSource: true,
		Level:      slogLogLevel,
		NoColor:    env.Active == env.PRO,
	})))

	// 初始化机器人
	bot := bot.NewBot()

	// 运行机器人
	bot.Run()
}
