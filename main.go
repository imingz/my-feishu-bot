package main

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/wsclient"
)

func main() {
	err := wsclient.Get().Start(context.Background())
	if err != nil {
		panic(err)
	}
}
