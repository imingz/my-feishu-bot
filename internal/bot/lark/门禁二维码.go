package lark

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/skip2/go-qrcode"
)

func (c *Client) Send门禁二维码消息卡片(ctx context.Context) error {
	// 1. 从 ctx 获取 messageId
	event := ctx.Value(consts.KeyEvent).(*larkim.P2MessageReceiveV1Data)
	messageId := event.Message.MessageId
	if messageId == nil {
		return fmt.Errorf("messageId 为空")
	}
	slog.Debug("", "messageId", *messageId)

	// 2. 生成卡片消息
	card, err := c.Generate门禁二维码消息卡片(*messageId)
	if err != nil {
		return err
	}

	return c.sendQrcodeCard(*messageId, card)
}

func (c *Client) Generate门禁二维码消息卡片(messageId string) (*larkcard.MessageCard, error) {
	// 1. 获取二维码图片
	gotQrcode, err := c.Huihutong.GetQrcodeData()
	if err != nil {
		return nil, fmt.Errorf("获取二维码图片失败, err: %v", err)
	}
	content, err := qrcode.Encode(gotQrcode, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("生成二维码图片失败, err: %v", err)
	}
	// 2. 上传图片
	imageKey, err := c.uploadImage(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("上传图片失败, err: %v", err)
	}
	slog.Debug("上传图片成功", "imageKey", imageKey)
	// 3. 获取消息卡片
	card := getQrcodeCard(imageKey)
	cardStr, _ := card.String()
	slog.Debug("生成消息卡片成功", slog.String("card", cardStr))
	return card, nil
}

// 上传图片，返回图片key
func (c *Client) uploadImage(image io.Reader) (string, error) {
	// 构造请求
	req := larkim.NewCreateImageReqBuilder().
		Body(larkim.NewCreateImageReqBodyBuilder().
			ImageType("message").
			Image(image).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Im.Image.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", err
	}

	// 服务端错误处理
	if !resp.Success() {
		return "", fmt.Errorf("client.Im.Image.Create failed, code: %d, msg: %s, log_id: %s", resp.Code, resp.Msg, resp.RequestId())
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
func (c *Client) sendQrcodeCard(messageId string, card *larkcard.MessageCard) error {
	// 创建请求对象
	cardJson, err := card.JSON()
	if err != nil {
		return err
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
	resp, err := c.client.Im.Message.Reply(context.Background(), req)

	// 处理错误
	if err != nil {
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		return fmt.Errorf("client.Im.Message.Reply failed, code: %d, msg: %s, log_id: %s", resp.Code, resp.Msg, resp.RequestId())
	}

	return nil
}
