package lark

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"xiaoxiaojiqiren/config"
	huihutongclient "xiaoxiaojiqiren/internal/pkg/HuiHutong_client"
	"xiaoxiaojiqiren/internal/pkg/consts"
	"xiaoxiaojiqiren/internal/pkg/slogor"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func getClient() *Client {
	slog.SetDefault(slog.New(slogor.NewHandler(os.Stderr, slogor.Options{
		TimeFormat: "2006-01-02 15:04:05.000",
		ShowSource: true,
		Level:      slog.LevelDebug,
	})))

	config := config.NewConfig()

	return NewClient(config.Bot.AppID, config.Bot.AppSecret, huihutongclient.NewClient())
}

func TestClient_SendRoomBalanceText(t *testing.T) {
	c := getClient()
	open_id := "ou_3d3a97494b21e6b6cb60e58ee8e39e00"
	c.SendRoomBalanceText(context.WithValue(context.Background(), consts.KeyEvent, &larkim.P2MessageReceiveV1Data{
		Sender: &larkim.EventSender{
			SenderId: &larkim.UserId{
				OpenId: &open_id,
			},
		},
	}))
}
