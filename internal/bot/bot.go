package bot

import (
	"context"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/env"
	"xiaoxiaojiqiren/internal/pkg/handler"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

type Bot struct {
	Config     *config.Config // 配置
	WsClient   *larkws.Client // 事件订阅客户端
	HttpClient *gin.Engine    // HTTP 服务
}

func NewBot() *Bot {
	// 初始化配置
	config := config.NewConfig()

	// 初始化 HTTP 服务
	httpClient := gin.Default()
	httpClient.POST("/webhook/card", sdkginext.NewCardActionHandlerFunc(cardHandler))

	// 根据环境设置日志级别
	logLevel := larkcore.LogLevelDebug
	if env.Active == env.PRO {
		logLevel = larkcore.LogLevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化 WebSocket 客户端
	var wsClient *larkws.Client
	wsHandler := dispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(P2MessageReceive)

	wsClient = larkws.NewClient(config.APP.ID, config.APP.Secret,
		larkws.WithEventHandler(wsHandler),
		larkws.WithLogLevel(logLevel),
	)

	return &Bot{
		Config:     config,
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
