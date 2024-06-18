package handler

import (
	"context"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/qrcode"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 接受到任意消息就回复二维码
func P2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	if event.Event.Message.ThreadId != nil {

		if *event.Event.Message.ThreadId == config.Get().Qrcode.ThreadId {
			img, err := qrcode.GetQrcodeFile()
			if err != nil {
				return err
			}
			imageKey, err := client.Get().Im_Image_Upload(client.ImageType_Message, img)
			if err != nil {
				return err
			}
			return client.Get().Im_Message_Reply("om_b0d3002db89790f29b1c2f96fd7881fe", client.Im_Message_Reply_Request{
				Content: `{
			"image_key": "` + imageKey + `"
		}`,
				MsgType:       "image",
				ReplyInThread: true,
				Uuid:          "",
			})
		}
	}
	return nil
}
