package handler

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 接受到任意消息就回复二维码
func P2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	if event.Event.Message.ThreadId != nil {
		switch *event.Event.Message.ThreadId {
		case config.Get().Qrcode.ThreadId:
			return biz.SendQrcodeCard(*event.Event.Message.MessageId)
		}
	}
	return nil
}
