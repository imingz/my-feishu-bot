package ws

import (
	"context"
	"fmt"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// p2MessageReceive 处理 p2MessageReceive 事件
func p2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	fmt.Println(larkcore.Prettify(event))

	switch *event.Event.Message.MessageType {
	case "text":
		text := getText(*event.Event.Message.Content)

		handler, ok := keyWords2Handler[text]
		if !ok {
			return nil
		}

		slog.Info("收到关键词消息", "关键词", text, "Sender.OpenId", *event.Event.Sender.SenderId.OpenId)
		ctx = context.WithValue(ctx, consts.KeyEvent, event.Event)
		if err := handler(ctx); err != nil {
			slog.Error("处理关键词消息失败", "关键词", text, "err", err)
			return err
		}
	}
	return nil
}
