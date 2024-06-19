package handler

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
)

var CardHandler = larkcard.NewCardActionHandler(config.Get().APP.VerificationToken, config.Get().APP.EventEncryptKey, qrcodeCardHandler)

func qrcodeCardHandler(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
	if cardAction.Action.Value["qrcode"] == "refresh" {
		return biz.GenerateQrcodeCard(cardAction.OpenMessageID)
	}
	return nil, nil
}
