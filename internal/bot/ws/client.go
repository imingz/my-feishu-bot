package ws

import (
	"xiaoxiaojiqiren/internal/bot/lark"
	"xiaoxiaojiqiren/internal/pkg/env"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

type Client struct {
	*larkws.Client
	LarkClient *lark.Client
}

func NewClient(appID, appSecret string, larkClient *lark.Client) *Client {
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
		Client:     wsClient,
		LarkClient: larkClient,
	}
}
