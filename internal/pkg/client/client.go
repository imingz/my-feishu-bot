package client

import (
	"sync"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/env"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type Client struct {
	*lark.Client
}

var instance *Client
var once sync.Once

func Get() *Client {
	once.Do(func() {
		switch env.Active {
		case env.DEV:
			instance = &Client{
				Client: lark.NewClient(config.Get().APP.ID, config.Get().APP.Secret, lark.WithLogLevel(larkcore.LogLevelDebug)),
			}
		case env.PRO:
			instance = &Client{
				Client: lark.NewClient(config.Get().APP.ID, config.Get().APP.Secret, lark.WithLogLevel(larkcore.LogLevelInfo)),
			}
		}

	})

	return instance
}
