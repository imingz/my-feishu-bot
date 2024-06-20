package ginclient

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
)

func (c *Client) cardHandler() gin.HandlerFunc {
	cardHandlerFunc := func(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
		if cardAction.Action.Value["qrcode"] == "refresh" {
			slog.Info("收到刷新二维码请求", "cardAction.OpenID", cardAction.OpenID)
			card, err := c.larkClient.GenerateQrcodeCard(cardAction.OpenMessageID)
			if err != nil {
				slog.Error("生成二维码卡片失败", "err", err)
			}
			return card, err
		}
		return nil, nil
	}

	cardHandler := larkcard.NewCardActionHandler(c.verificationToken, c.eventEncryptKey, cardHandlerFunc)

	return sdkginext.NewCardActionHandlerFunc(cardHandler)
}
