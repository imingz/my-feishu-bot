package qrcodeclient

import (
	"encoding/json"
	"fmt"
	"time"
	"xiaoxiaojiqiren/internal/pkg/config"

	"github.com/imroc/req/v3"
)

var client = req.C().SetBaseURL(config.Get().Qrcode.BaseUrl).
	SetTimeout(5 * time.Second).
	EnableDumpEachRequest()

type response struct {
	Code      int64           `json:"code"`
	Data      json.RawMessage `json:"data"`
	Message   string          `json:"message"`
	RequestID string          `json:"requestId"`
	Success   bool            `json:"success"`
	Timestamp int64           `json:"timestamp"`
}

func (resp response) getError() string {
	return fmt.Sprintf("code: %d, message: %s", resp.Code, resp.Message)
}
