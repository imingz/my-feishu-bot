package wsclient

import (
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/wsclient/handler"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/ws"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

var client = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
	larkws.WithEventHandler(handler.Get()),
	larkws.WithLogLevel(larkcore.LogLevelDebug),
)

func Get() *ws.Client {
	return client
}
