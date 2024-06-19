package handler

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/biz"
	"xiaoxiaojiqiren/internal/pkg/config"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 接受到任意消息就回复二维码
func P2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	switch *event.Event.Message.ChatId {
	case config.Get().Qrcode.ChatId: // 可以接受二维码的群
		if *event.Event.Message.MessageType == "text" {
			const CorrectText = `{"text":"门禁二维码"}`
			if *event.Event.Message.Content == CorrectText {
				// 回复消息为话题并发送初始二维码卡片
				return biz.SendQrcodeCard(*event.Event.Message.MessageId)
			}
		}
	}
	return nil
}
