package handler

import (
	"context"
	"fmt"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

var CardHandler = larkcard.NewCardActionHandler(config.Get().APP.VerificationToken, config.Get().APP.EventEncryptKey, qrcodeCardHandler)

func qrcodeCardHandler(ctx context.Context, cardAction *larkcard.CardAction) (any, error) {
	fmt.Println(larkcore.Prettify(cardAction))
	if cardAction.Action.Value["qrcode"] == "refresh" {
		err := biz.SendQrcodeCard(cardAction.OpenMessageID)
		return nil, err
	}
	return nil, nil
}
