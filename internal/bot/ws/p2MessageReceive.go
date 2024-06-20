package ws

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// p2MessageReceive 处理 p2MessageReceive 事件
func p2MessageReceive(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	fmt.Println(larkcore.Prettify(event))
	return nil
}
