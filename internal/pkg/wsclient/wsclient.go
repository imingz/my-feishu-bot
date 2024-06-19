package wsclient

import (
	"sync"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/wsclient/handler"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

var client *larkws.Client

var once sync.Once

func Get() *larkws.Client {
	once.Do(func() {
		switch env.Active {
		case env.DEV:
			client = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
				larkws.WithEventHandler(handler.Get()),
				larkws.WithLogLevel(larkcore.LogLevelDebug),
			)
		case env.PRO:
			client = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
				larkws.WithEventHandler(handler.Get()),
				larkws.WithLogLevel(larkcore.LogLevelInfo),
			)
		}
	})
	return client
}
