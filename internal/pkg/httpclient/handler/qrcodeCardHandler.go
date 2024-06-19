package handler

import (
	"context"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
)

var CardHandler = larkcard.NewCardActionHandler(config.Get().APP.VerificationToken, config.Get().APP.EventEncryptKey, qrcodeCardHandler)

func qrcodeCardHandler(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
	if cardAction.Action.Value["qrcode"] == "refresh" {
		slog.Info("收到刷新二维码请求", "cardAction.OpenID", cardAction.OpenID)
		card, err := biz.GenerateQrcodeCard(cardAction.OpenMessageID)
		if err != nil {
			slog.Error("生成二维码卡片失败", "err", err)
		}
		return card, err
	}
	return nil, nil
}
