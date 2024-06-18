package client

import (
	"xiaoxiaojiqiren/internal/pkg/config"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

var client = lark.NewClient(config.Get().APP.ID, config.Get().APP.Secret, lark.WithLogLevel(larkcore.LogLevelDebug))

func Get() *lark.Client {
	return client
}
