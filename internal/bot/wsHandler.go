package bot

import (
	"context"
	"encoding/json"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/consts"
	"xiaoxiaojiqiren/internal/pkg/handler"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

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
				slog.Error("处理关键词消息失败", "关键词", text.Text, "err", err)
				return err
			}
		}
	}
	return nil
}
