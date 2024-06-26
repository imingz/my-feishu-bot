package ws

import (
	"context"
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/pkg/env"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

type Client struct {
	*larkws.Client
}

var keyWords2Handler = make(map[string]func(ctx context.Context) error)

func NewClient(appID, appSecret string, larkClient *lark.Client) *Client {
	keyWords2Handler["电费"] = larkClient.Send宿舍电费余额文本
	keyWords2Handler["门禁二维码"] = larkClient.Send门禁二维码消息卡片

	wsHandler := dispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(p2MessageReceive)

	wsLogLevel := larkcore.LogLevelInfo
	if env.Active == env.DEV {
		wsLogLevel = larkcore.LogLevelDebug
	}

	wsClient := larkws.NewClient(appID, appSecret,
		larkws.WithEventHandler(wsHandler),
		larkws.WithLogLevel(wsLogLevel),
	)
	return &Client{
		Client: wsClient,
	}
}
