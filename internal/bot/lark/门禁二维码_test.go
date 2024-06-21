package lark

import (
	"context"
	"fmt"
	"testing"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func TestSendQrcodeCard(t *testing.T) {
	c := getClient()
	msg_id := "om_42b9a6ca276032a9673e5f825210773b"
	err := c.Send门禁二维码消息卡片(context.WithValue(context.Background(), consts.KeyEvent, &larkim.P2MessageReceiveV1Data{
		Sender: &larkim.EventSender{},
		Message: &larkim.EventMessage{
			MessageId: &msg_id,
		},
	}))
	fmt.Printf("err: %v\n", err)
}
