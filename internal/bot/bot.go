package bot

import (
	"context"
	"encoding/json"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/consts"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/handler"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

type Bot struct {
	WsClient   *larkws.Client
	HttpClient *gin.Engine
}

func NewBot() *Bot {
	httpClient := gin.Default()
	httpClient.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(cardHandler))

	// 根据环境设置日志级别
	logLevel := larkcore.LogLevelDebug
	if env.Active == env.PRO {
		logLevel = larkcore.LogLevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	var wsClient *larkws.Client
	wsHandler := dispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(P2MessageReceive)

	wsClient = larkws.NewClient(config.Get().APP.ID, config.Get().APP.Secret,
		larkws.WithEventHandler(wsHandler),
		larkws.WithLogLevel(logLevel),
	)

	return &Bot{
		WsClient:   wsClient,
		HttpClient: httpClient,
	}
}

func (b *Bot) Run() {
	go func() {
		b.HttpClient.Run()
	}()

	err := b.WsClient.Start(context.Background())
	if err != nil {
		panic(err)
	}
}

type HandlerFunc func(ctx context.Context) error

var keyWords2Handler map[string]HandlerFunc = map[string]HandlerFunc{
	"门禁二维码": handler.SendQrcodeCard,
	"宿舍电费":  handler.SendRoomBalanceText,
}

// P2MessageReceive 处理 P2MessageReceive 事件
func P2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	switch *event.Event.Message.ChatId {
	case config.Get().Qrcode.ChatId: // 可以接受二维码的群 	// TODO: 优化结构，这里不只有二维码了，还可以查询电费
		switch *event.Event.Message.MessageType {
		case "text":
			var text struct {
				Text string `json:"text"`
			}
			if err := json.Unmarshal([]byte(*event.Event.Message.Content), &text); err != nil {
				slog.Error("解析消息内容失败", "err", err)
				return err
			}

			handler, ok := keyWords2Handler[text.Text]
			if !ok {
				return nil
			}

			slog.Info("收到关键词消息", "关键词", text.Text, "Sender.OpenId", *event.Event.Sender.SenderId.OpenId)
			ctx = context.WithValue(ctx, consts.KeyMessageID, *event.Event.Message.MessageId)
			if err := handler(ctx); err != nil {
				slog.Error("生成二维码卡片消息失败", "err", err)
				return err
			}
		}
	}
	return nil
}

var cardHandler = larkcard.NewCardActionHandler(config.Get().APP.VerificationToken, config.Get().APP.EventEncryptKey, qrcodeCardHandler)

func qrcodeCardHandler(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
	if cardAction.Action.Value["qrcode"] == "refresh" {
		slog.Info("收到刷新二维码请求", "cardAction.OpenID", cardAction.OpenID)
		card, err := handler.GenerateQrcodeCard(cardAction.OpenMessageID)
		if err != nil {
			slog.Error("生成二维码卡片失败", "err", err)
		}
		return card, err
	}
	return nil, nil
}
