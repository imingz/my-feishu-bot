package biz

import (
	"context"
	"fmt"
	"xiaoxiaojiqiren/internal/pkg/client"
	"xiaoxiaojiqiren/internal/pkg/config"
	"xiaoxiaojiqiren/internal/pkg/hhtclient"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func SendRoomBalanceText() error {
	// 创建请求对象
	balance, err := hhtclient.GetRoomBalance(context.Background())
	if err != nil {
		return err
	}
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
			ReceiveId(config.Get().RoomBalance.OpenId).
			MsgType(`post`).
			Content(content).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Get().Im.Message.Create(context.Background(), req)

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
