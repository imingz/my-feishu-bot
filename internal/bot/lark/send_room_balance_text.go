package lark

import (
	"context"
	"fmt"
	"log/slog"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func (c *Client) SendRoomBalanceText(ctx context.Context) error {
	// 获取 messageId
	event := ctx.Value(consts.KeyEvent).(*larkim.P2MessageReceiveV1Data)
	messageId := event.Message.MessageId
	if messageId == nil {
		return fmt.Errorf("message_id 为空")
	}

	// 获取房间配置
	info, err := c.获取宿舍电费房间信息(ctx)
	if err != nil {
		return fmt.Errorf("获取房间配置失败: %v", err)
	}

	// 获取电费余额
	balance, err := c.huihutong.GetRoomBalance(info.楼栋号, info.楼层, info.房间号)
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

	req := larkim.NewReplyMessageReqBuilder().
		MessageId(*messageId).
		Body(larkim.NewReplyMessageReqBodyBuilder().
			Content(content).
			MsgType(`post`).
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
		return fmt.Errorf("resp.Code: %v, resp.Msg: %v, resp.RequestId(): %v", resp.Code, resp.Msg, resp.RequestId())
	}

	return nil
}
