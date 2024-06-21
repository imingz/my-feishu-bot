package cron

import (
	"log/slog"
	"time"
	"xiaoxiaojiqiren/internal/bot/lark"

	"github.com/robfig/cron/v3"
)

type Client struct {
	cron       *cron.Cron // 定时任务
	larkClient *lark.Client
}

func New(larkClient *lark.Client) *Client {
	c := &Client{
		cron:       cron.New(),
		larkClient: larkClient,
	}

	// 维护定时任务
	var taskMap = map[string]func(){
		"0 20 * * *": c.检查宿舍电费, // 每天 20 点检查宿舍电费
		"0 * * * *":  c.检查宿舍电费,
	}

	// 添加定时任务
	for spec, fn := range taskMap {

		id, err := c.cron.AddFunc(spec, fn)
		if err != nil {
			slog.Error("添加定时任务失败", "spec", spec, "任务名称", getFunctionName(fn), "错误", err)
			continue
		}

		// 解析 Cron 表达式
		schedule, err := cron.ParseStandard(spec)
		if err != nil {
			slog.Error("解析 Cron 表达式失败:", err)
		}

		slog.Info("已添加定时任务", "spec", spec, "任务名称", getFunctionName(fn), "ID", id, "下次执行时间", schedule.Next(time.Now()))
	}

	return c
}

func (c *Client) Start() {
	c.cron.Start()
}
