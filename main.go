package main

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/httpclient"
	"xiaoxiaojiqiren/internal/pkg/wsclient"
)

func main() {
	// 启动http服务
	go func() {
		httpclient.Get().Run()
	}()

	err := wsclient.Get().Start(context.Background())
	if err != nil {
		panic(err)
	}
}
