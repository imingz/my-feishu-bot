package cron

import "github.com/robfig/cron/v3"

type Client struct {
	cron *cron.Cron // 定时任务
}

func New() *Client {
	return &Client{
		cron: cron.New(),
	}
}

func (c *Client) Start() {
	c.cron.Start()
}
