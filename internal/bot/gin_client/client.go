package ginclient

import (
	"xiaoxiaojiqiren/internal/bot/lark"

	"github.com/gin-gonic/gin"
)

type Client struct {
	*gin.Engine
	verificationToken string
	eventEncryptKey   string
	larkClient        *lark.Client
}

func NewClient(verificationToken, eventEncryptKey string, larkClient *lark.Client) *Client {
	g := gin.Default()

	c := &Client{
		Engine:            g,
		verificationToken: verificationToken,
		eventEncryptKey:   eventEncryptKey,
		larkClient:        larkClient,
	}

	webhook := g.Group("/webhook")
	{
		webhook.POST("/lark", c.cardHandler())
		webhook.POST("/17track", c.webhook17trackHandler)
	}

	return c
}
