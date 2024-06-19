package reqclient

import (
	"time"
	"xiaoxiaojiqiren/internal/pkg/config"

	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
}

func NewClient(config *config.Config) *Client {
	return &Client{
		Client: req.C().SetBaseURL(config.Qrcode.BaseUrl).
			SetTimeout(3 * time.Second).
			EnableDumpEachRequest(),
	}
}
