package lark

import (
	"context"
	"fmt"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func (c *Client) SendRoomBalanceText(ctx context.Context) error {
	// 获取接收者 open_id
	event := ctx.Value(consts.KeyEvent).(*larkim.P2MessageReceiveV1Data)
	open_id := event.Sender.SenderId.OpenId
	if open_id == nil {
		return fmt.Errorf("open_id 为空")
	}

	// 获取电费余额
	balance, err := c.huihutong.GetRoomBalance()
	if err != nil {
		return err
	}
	slog.Debug("电费余额", "balance", balance)

	// 创建请求对象
	content, err := larkim.NewMessagePost().
		ZhCn(larkim.NewMessagePostContent().
			ContentTitle(`电费余额`).
			AppendContent([]larkim.MessagePostElement{&larkim.MessagePostText{Text: fmt.Sprintf("%.2f", balance)}}).
			Build()).
		Build()
	if err != nil {
		return err
	}
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`open_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(*open_id).
			MsgType(`post`).
			Content(content).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		return fmt.Errorf("resp.Code: %v, resp.Msg: %v, resp.RequestId(): %v", resp.Code, resp.Msg, resp.RequestId())
	}

	return nil
}
