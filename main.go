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
		NoColor:    env.Active == env.PRO,
	})))

	bot := bot.NewBot()
	bot.Run()
}
