package wsclient

import (
	"sync"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/env"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

var client *larkws.Client
var once sync.Once

func Get() *larkws.Client {
	once.Do(func() {
		handler := dispatcher.NewEventDispatcher("", "").
			OnP2MessageReceiveV1(P2MessageReceive)
		switch env.Active {
		case env.DEV:
			client = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
				larkws.WithEventHandler(handler),
				larkws.WithLogLevel(larkcore.LogLevelDebug),
			)
		case env.PRO:
			client = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
				larkws.WithEventHandler(handler),
				larkws.WithLogLevel(larkcore.LogLevelInfo),
			)
		}
	})
	return client
}
