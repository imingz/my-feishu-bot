package client

import (
	"context"
	"fmt"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type Im_Message_Reply_Request struct {
	Content       string // 消息内容 json 格式
	MsgType       string // 消息类型，包括：text、post、image、file、audio、media、sticker、interactive、share_card、share_user
	ReplyInThread bool   // 是否以话题形式回复；若要回复的消息已经是话题消息，则默认以话题形式进行回复
	Uuid          string
}

func (c *Client) Im_Message_Reply(message_id string, request Im_Message_Reply_Request) error {
	// 创建请求对象
	req := larkim.NewReplyMessageReqBuilder().
		MessageId(message_id).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(request.Content).
			MsgType(request.MsgType).
			ReplyInThread(request.ReplyInThread).
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
