package cron

import (
	"fmt"
	"log/slog"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

func (c *Client) 检查宿舍电费() {
	slog.Info("启动定时任务", "任务", "检查宿舍电费")

	slog.Debug("获取宿舍电费配置...")
	items, err := c.larkClient.Get宿舍电费配置()
	if err != nil {
		slog.Error("获取宿舍电费配置失败", "err", err)
	}
	slog.Debug("获取宿舍电费配置成功", "items", larkcore.Prettify(items))

	for _, item := range items {
		slog.Info("正在处理",
			slog.String("name", item.Fields["提交人"].([]any)[0].(map[string]any)["name"].(string)),
			slog.String("open_id", item.Fields["提交人"].([]any)[0].(map[string]any)["id"].(string)),
		)
		// 获取企业邮箱
		enterpriseEmail, err := c.larkClient.GetEnterpriseEmailByOpenId(item.Fields["提交人"].([]any)[0].(map[string]any)["id"].(string))
		if err != nil {
			slog.Error("获取企业邮箱失败", "err", err)
			continue
		}
		slog.Debug("获取企业邮箱成功", "enterpriseEmail", enterpriseEmail)

		// 获取 楼栋号，楼层，房间号
		楼栋号 := fmt.Sprintf("%.0f", item.Fields["宿舍楼栋号"].(float64))
		楼层 := fmt.Sprintf("%.0f", item.Fields["楼层"].(float64))
		房间号 := fmt.Sprintf("%02.0f", item.Fields["房间号"].(float64))

		// 获取电费余额
		balance, err := c.larkClient.Huihutong.Get房间余额(楼栋号, 楼层, 房间号)
		if err != nil {
			slog.Error("获取电费余额失败", "err", err)
			continue
		}

		// 获取电费阈值
		electricityThreshold := item.Fields["电费阈值"].(float64)
		slog.Info("获取电费余额成功", "balance", balance, "electricityThreshold", electricityThreshold)

		if balance <= electricityThreshold {
			// 发送邮件
			slog.Debug("电费不足，发送邮件")
			err := c.larkClient.Mail.Send电费不足(enterpriseEmail, balance)
			if err != nil {
				slog.Error("发送邮件失败", "err", err)
				continue
			}
			slog.Info("发送邮件成功")
		}
	}

}
