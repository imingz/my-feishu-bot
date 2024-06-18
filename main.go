package main

import (
	"context"
	"net/http"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/wsclient"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
)

func main() {
	go func() {
		// 创建 card 处理器
		cardHandler := larkcard.NewCardActionHandler(config.Get().APP.VerificationToken, config.Get().APP.EventEncryptKey, func(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
			if cardAction.Action.Value["qrcode"] == "refresh" {
				biz.SendQrcodeCard(cardAction.OpenMessageID)
			}
			return nil, nil
		})

		// 注册处理器
		http.HandleFunc("/card", httpserverext.NewCardActionHandlerFunc(cardHandler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

		// 启动http服务
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	err := wsclient.Get().Start(context.Background())
	if err != nil {
		panic(err)
	}
}
