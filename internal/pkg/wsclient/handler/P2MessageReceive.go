package handler

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/qrcode"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 接受到任意消息就回复二维码
func P2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	img, err := qrcode.GetQrcodeFile()
	if err != nil {
		return err
	}
	_, err = client.SendImage(client.Get(), &client.SendImageRequest{
		Image:         img,
		ReceiveIdType: "open_id",
		ReceiveId:     *event.Event.Sender.SenderId.OpenId,
	})
	return err
}
