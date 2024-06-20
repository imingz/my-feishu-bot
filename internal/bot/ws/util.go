package ws

import (
	"encoding/json"
	"log/slog"
)

func getText(data string) string {
	var text struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal([]byte(data), &text); err != nil {
		slog.Error("解析消息内容失败", "err", err)
		return ""
	}
	return text.Text
}
