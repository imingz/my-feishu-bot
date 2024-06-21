package lark

import (
	"log/slog"
	"testing"
)

func TestClient_Get企业邮箱ByOpenId(t *testing.T) {
	c := getClient()
	EnterpriseEmail, err := c.GetEnterpriseEmailByOpenId(`ou_3d3a97494b21e6b6cb60e58ee8e39e00`)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info(EnterpriseEmail)
}
