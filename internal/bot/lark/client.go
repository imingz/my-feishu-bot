package lark

import (
	"xiaoxiaojiqiren/internal/pkg/env"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type Client struct {
	*lark.Client
}

func NewClient(appID, appSecret string) *Client {
	logLevel := larkcore.LogLevelInfo
	if env.Active == env.DEV {
		logLevel = larkcore.LogLevelDebug
	}

	return &Client{lark.NewClient(appID, appSecret, lark.WithLogLevel(logLevel))}
}
