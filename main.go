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
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: "2006-01-02 15:04:05.000",
		ShowSource: true,
		Level:      slog.LevelDebug, // TODO: 根据环境设置日志级别
		NoColor:    env.Active == env.PRO,
	})))

	// 初始化机器人
	bot := bot.NewBot()

	// 运行机器人
	bot.Run()
}
