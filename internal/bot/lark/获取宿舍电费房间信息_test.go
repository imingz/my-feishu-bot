package lark

import (
	"context"
	"fmt"
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
	fmt.Printf("info: %+v\n", info)
}
