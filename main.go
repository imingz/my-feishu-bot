package main

import (
	"context"
	"log/slog"
	"os"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/httpclient"
	"xiaoxiaojiqiren/internal/pkg/slogor"
	"xiaoxiaojiqiren/internal/pkg/wsclient"
)

func main() {
	// 设置 slog
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: "2006-01-02 15:04:05.000",
		ShowSource: true,
		NoColor:    env.Active == env.PRO,
	})))

	// 启动http服务
	go func() {
		httpclient.Get().Run()
	}()

	err := wsclient.Get().Start(context.Background())
	if err != nil {
		panic(err)
	}
}
