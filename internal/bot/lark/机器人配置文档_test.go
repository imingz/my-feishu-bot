package lark

import (
	"log/slog"
	"testing"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

func TestClient_Get宿舍电费配置(t *testing.T) {
	c := getClient()
	items, err := c.Get宿舍电费配置()
	if err != nil {
		t.Fatal(err)
	}
	slog.Info(larkcore.Prettify(items))
}
