package ginclient

import (
	"log/slog"
	"xiaoxiaojiqiren/internal/bot/lark"

	"github.com/gin-gonic/gin"
)

// webhook17trackHandler 处理Webhook请求
func (client *Client) webhook17trackHandler(c *gin.Context) {
	var event _17trackEvent

	// 绑定数据
	if err := c.ShouldBindJSON(&event); err != nil {
		return
	}

	// 业务处理
	if event.Event == TRACKING_UPDATED {
		slog.Debug("TRACKING UPDATED",
			"物流单号", event.Data.Number,
			"运输商代码", event.Data.Carrier,
			"运输商名称", event.Data.TrackInfo.Tracking.Providers[0].Provider.Alias,
			"物流状态", _17trackStatus[event.Data.TrackInfo.LatestStatus.Status],
			"物流子状态", _17trackSubStatus[event.Data.TrackInfo.LatestStatus.SubStatus],
			"事件描述", event.Data.TrackInfo.LatestEvent.Description,
			"地点信息", event.Data.TrackInfo.LatestEvent.Location,
		)
	} else {
		slog.Debug("TRACKING STOPPED",
			"停止跟踪的物流单号", event.Data.Number,
			"停止跟踪的物流商", event.Data.Carrier,
			"运输商名称", event.Data.TrackInfo.Tracking.Providers[0].Provider.Alias,
		)
	}

	// 发送消息到指定群聊
	client.larkClient.Send快递状态(
		lark.DeliveryStatus{
			Number:      event.Data.Number,
			Carrier:     event.Data.TrackInfo.Tracking.Providers[0].Provider.Alias,
			Status:      _17trackStatus[event.Data.TrackInfo.LatestStatus.Status],
			SubStatus:   _17trackSubStatus[event.Data.TrackInfo.LatestStatus.SubStatus],
			Description: event.Data.TrackInfo.LatestEvent.Description,
			Location:    event.Data.TrackInfo.LatestEvent.Location,
		},
	)
}
