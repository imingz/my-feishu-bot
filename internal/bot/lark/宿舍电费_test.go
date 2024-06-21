package lark

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func TestClient_获取宿舍电费房间信息(t *testing.T) {
	c := getClient()
	open_id := `ou_3d3a97494b21e6b6cb60e58ee8e39e00`
	info, err := c.获取宿舍电费房间信息(context.WithValue(context.Background(), consts.KeyEvent, &larkim.P2MessageReceiveV1Data{
		Sender: &larkim.EventSender{
			SenderId: &larkim.UserId{
				OpenId: &open_id,
			},
		},
	},
	))
	if err != nil {
		t.Fatal(err)
	}
	slog.Info(fmt.Sprintf("info: %+v\n", info))
}

func TestClient_Send宿舍电费余额文本(t *testing.T) {
	c := getClient()
	open_id := `ou_3d3a97494b21e6b6cb60e58ee8e39e00`
	msg_id := `om_42b9a6ca276032a9673e5f825210773b`
	slog.Info("测试开始")
	err := c.Send宿舍电费余额文本(context.WithValue(context.Background(), consts.KeyEvent, &larkim.P2MessageReceiveV1Data{
		Sender: &larkim.EventSender{
			SenderId: &larkim.UserId{
				OpenId: &open_id,
			},
		},
		Message: &larkim.EventMessage{
			MessageId: &msg_id,
		},
	}))
	if err != nil {
		slog.Error("测试失败", "err", err)
		t.Fatal(err)
	}
	slog.Info("测试结束")
}
