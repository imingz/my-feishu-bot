package biz

import (
	"context"
	"time"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/qrcode"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func SendQrcodeCard(messageId string) error {
	// 再发送一次
	img, err := qrcode.GetQrcodeFile()
	if err != nil {
		return err
	}
	imageKey, err := client.Get().Im_Image_Upload(client.ImageType_Message, img)
	if err != nil {
		return err
	}
	// 创建请求对象
	req := larkim.NewReplyMessageReqBuilder().
		MessageId(messageId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(`{"type": "template", "data": { "template_id": "` + config.Get().Qrcode.CardId + `", "template_variable": {"time": "` + time.Now().Format(time.DateTime) + `", "imgKey": "` + imageKey + `"} } }`).
			MsgType(`interactive`).
			ReplyInThread(true).
			Build()).
		Build()
	_, err = client.Get().Im.Message.Reply(context.Background(), req)
	return err
}
