package lark

import (
	"context"
	"testing"
	"xiaoxiaojiqiren/config"
	huihutongclient "xiaoxiaojiqiren/internal/pkg/HuiHutong_client"
	"xiaoxiaojiqiren/internal/pkg/consts"
)

func getClient() *Client {
	config := config.NewConfig()

	return NewClient(config.Bot.AppID, config.Bot.AppSecret, huihutongclient.NewClient())
}

func TestClient_SendRoomBalanceText(t *testing.T) {
	c := getClient()
	c.SendRoomBalanceText(context.WithValue(context.Background(), consts.KeyOpenID, "ou_3d3a97494b21e6b6cb60e58ee8e39e00"))
}
