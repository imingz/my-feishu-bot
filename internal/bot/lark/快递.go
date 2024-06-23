package lark

import (
	"context"
	"fmt"
	"log/slog"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type DeliveryStatus struct {
	Number      string `json:"number"`
	Carrier     string `json:"carrier"`
	Status      string `json:"status"`
	SubStatus   string `json:"subStatus"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func (c *Client) Send快递状态(status DeliveryStatus) {
	slog.Info("发送快递状态")
	// 创建请求对象
	content := fmt.Sprintf("{\"type\": \"template\", \"data\": { \"template_id\": \"AAqH7hkOjpkYp\", \"template_variable\": {\"number\": \"%s\",\"carrier\": \"%s\",\"status\": \"%s\",\"subStatus\": \"%s\",\"description\":\"%s\", \"location\":\"%s\",\"eventType\": \"\"} } }",
		status.Number,
		status.Carrier,
		status.Status,
		status.SubStatus,
		status.Description,
		status.Location)
	slog.Debug(content)
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`chat_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(`oc_e2a0e15d49008d2745f09c2f86ec9968`). // TODO: 修改为特定的群聊ID
			MsgType(`interactive`).
			Content(content).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
