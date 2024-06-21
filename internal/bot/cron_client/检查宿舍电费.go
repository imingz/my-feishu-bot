package cron

import "log/slog"

func (c *Client) 检查宿舍电费() {
	slog.Info("启动定时任务", "任务", "定时检查宿舍电费")
}
