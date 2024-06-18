package biz

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/qrcode"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func SendQrcodeCard(messageId string) error {
	// 1. 获取二维码图片
	img, err := qrcode.GetQrcodeFile()
	if err != nil {
		return errors.Join(err, errors.New("获取二维码图片失败"))
	}
	// 2. 上传图片
	imageKey, err := uploadImage(img)
	if err != nil {
		return err
	}
	// 3. 发送卡片消息
	return sendQrcodeCard(messageId, imageKey)
}

// 上传图片，返回图片key
func uploadImage(image io.Reader) (string, error) {
	// 错误预定义
	var uploadImageErr = errors.New("上传图片失败")
	// 构造请求
	req := larkim.NewCreateImageReqBuilder().
		Body(larkim.NewCreateImageReqBodyBuilder().
			ImageType("message").
			Image(image).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Get().Im.Image.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", errors.Join(err, uploadImageErr)
	}

	// 服务端错误处理
	if !resp.Success() {
		return "", errors.Join(fmt.Errorf("client.Im.Image.Create failed, code: %d, msg: %s, log_id: %s", resp.Code, resp.Msg, resp.RequestId()), uploadImageErr)
	}

	return *resp.Data.ImageKey, nil
}

// 发送二维码卡片消息
func sendQrcodeCard(messageId, imageKey string) error {
	// 错误预定义
	var sendQrcodeCardErr = errors.New("发送二维码卡片消息失败")
	// 创建请求对象
	req := larkim.NewReplyMessageReqBuilder().
		MessageId(messageId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(`{"type": "template", "data": { "template_id": "` + config.Get().Qrcode.CardId + `", "template_variable": {"time": "` + time.Now().Format(time.DateTime) + `", "imgKey": "` + imageKey + `"} } }`).
			MsgType(`interactive`).
			ReplyInThread(true).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Get().Im.Message.Reply(context.Background(), req)

	// 处理错误
	if err != nil {
		return errors.Join(err, sendQrcodeCardErr)
	}

	// 服务端错误处理
	if !resp.Success() {
		return errors.Join(fmt.Errorf("client.Im.Message.Reply failed, code: %d, msg: %s, log_id: %s", resp.Code, resp.Msg, resp.RequestId()), sendQrcodeCardErr)
	}

	return nil
}
