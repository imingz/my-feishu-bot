package biz

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/qrcode"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func GenerateQrcodeCard(messageId string) (*larkcard.MessageCard, error) {
	// 1. 获取二维码图片
	img, err := qrcode.GetQrcodeFile()
	if err != nil {
		return nil, err
	}
	// 2. 上传图片
	imageKey, err := uploadImage(img)
	if err != nil {
		slog.Error("上传图片失败", "err", err)
		return nil, err
	}
	// 3. 获取消息卡片
	return getQrcodeCard(imageKey), nil
}

func SendQrcodeCard(messageId string) error {
	// 1. 生成卡片消息
	card, err := GenerateQrcodeCard(messageId)
	if err != nil {
		return err
	}
	// 2. 发送卡片消息
	return sendQrcodeCard(messageId, card)
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

// 获取卡片
func getQrcodeCard(imageKey string) *larkcard.MessageCard {
	// header
	header := larkcard.NewMessageCardHeader().
		Template(larkcard.TemplateBlue).
		Title(larkcard.NewMessageCardPlainText().
			Content(time.Now().Format(time.DateTime)).
			Build()).
		Build()

	// 二维码图片
	qrcodeElement := larkcard.NewMessageCardImage().
		ImgKey(imageKey).
		Alt(larkcard.NewMessageCardPlainText().Content("图片").Build()).
		CompactWidth(true).
		Build()

	// 重新生成按钮
	generateButton := larkcard.NewMessageCardAction().
		Actions([]larkcard.MessageCardActionElement{
			larkcard.NewMessageCardEmbedButton().
				Type(larkcard.MessageCardButtonTypePrimary).
				Text(larkcard.NewMessageCardPlainText().
					Content("重新生成").
					Build()).
				Value(map[string]any{
					"qrcode": "refresh",
				}).
				Build(),
		}).
		Build()

	// 卡片消息体
	messageCard := larkcard.NewMessageCard().
		Header(header).
		Elements([]larkcard.MessageCardElement{qrcodeElement, generateButton}).
		Build()

	return messageCard
}

// 发送二维码卡片消息
func sendQrcodeCard(messageId string, card *larkcard.MessageCard) error {
	// 错误预定义
	var sendQrcodeCardErr = errors.New("发送二维码卡片消息失败")
	// 创建请求对象
	cardJson, err := card.JSON()
	if err != nil {
		return errors.Join(err, sendQrcodeCardErr)
	}
	req := larkim.NewReplyMessageReqBuilder().
		MessageId(messageId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(cardJson).
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
