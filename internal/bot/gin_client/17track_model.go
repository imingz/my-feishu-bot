package ginclient

type _17trackEventType string

const (
	// 自动跟踪后若物流信息有变化，则会即刻触发推送，反之不会；
	TRACKING_UPDATED = "TRACKING_UPDATED"
	// 收到此通知表示当前物流单号停止更新了，也不会有更新跟踪通知；
	// 如果单号停止跟踪是由系统规则触发的，则会推送停止跟踪消息；
	// 可以通过 重启跟踪 接口恢复自动跟踪；
	TRACKING_STOPPED = "TRACKING_STOPPED"
)

var _17trackStatus = map[string]string{
	"NotFound":           "查询不到，进行查询操作但没有得到结果，原因请参看子状态。",
	"InfoReceived":       "收到信息，运输商收到下单信息，等待上门取件。",
	"InTransit":          "运输途中，包裹正在运输途中，具体情况请参看子状态。",
	"Expired":            "运输过久，包裹已经运输了很长时间而仍未投递成功。",
	"AvailableForPickup": "到达待取，包裹已经到达目的地的投递点，需要收件人自取。",
	"OutForDelivery":     "派送途中，包裹正在投递过程中。",
	"DeliveryFailure":    "投递失败，包裹尝试派送但未能成功交付，原因请参看子状态。原因可能是：派送时收件人不在家、投递延误重新安排派送、收件人要求延迟派送、地址不详无法派送、因偏远地区不提供派送服务等。",
	"Delivered":          "成功签收，包裹已妥投。",
	"Exception":          "可能异常，包裹可能被退回，原因请参看子状态。原因可能是：收件人地址错误或不详、收件人拒收、包裹无人认领超过保留期等。包裹可能被海关扣留，常见扣关原因是：包含敏感违禁、限制进出口的物品、未交税款等。包裹可能在运输途中遭受损坏、丢失、延误投递等特殊情况。",
}

var _17trackSubStatus = map[string]string{
	"NotFound_Other":                        "运输商没有返回信息。",
	"NotFound_InvalidCode":                  "物流单号无效，无法进行查询。",
	"InfoReceived":                          "收到信息，暂无细分含义与主状态一致。",
	"InTransit_PickedUp":                    "已揽收，运输商已从发件人处取回包裹。",
	"InTransit_Other":                       "其它情况，暂无细分除当前已知子状态之外的情况。",
	"InTransit_Departure":                   "已离港，货物离开起运地（国家/地区）港口。",
	"InTransit_Arrival":                     "已到港，货物到达目的地（国家/地区）港口。",
	"InTransit_CustomsProcessing":           "清关中，货物在海关办理进入或出口的相关流程中。",
	"InTransit_CustomsReleased":             "清关完成，货物在海关完成了进入或出口的流程。",
	"InTransit_CustomsRequiringInformation": "需要资料，在清关中流程中承运人需要提供相关资料才能完成清关。",
	"Expired_Other":                         "运输过久，暂无细分含义与主状态一致。",
	"AvailableForPickup_Other":              "到达待取，暂无细分含义与主状态一致。",
	"OutForDelivery_Other":                  "派送途中，暂无细分含义与主状态一致。",
	"DeliveryFailure_Other":                 "其它情况，暂无细分除当前已知子状态之外的情况。",
	"DeliveryFailure_NoBody":                "找不到收件人，派送中的包裹暂时无法联系上收件人，导致投递失败。",
	"DeliveryFailure_Security":              "安全原因，派送中发现的包裹安全、清关、费用问题，导致投递失败。",
	"DeliveryFailure_Rejected":              "拒收，收件人因某些原因拒绝接收包裹，导致投递失败。",
	"DeliveryFailure_InvalidAddress":        "地址错误，由于收件人地址不正确，导致投递失败。",
	"Delivered_Other":                       "成功签收，暂无细分含义与主状态一致。",
	"Exception_Other":                       "其它情况，暂无细分除当前已知子状态之外的情况。",
	"Exception_Returning":                   "退件中，包裹正在送回寄件人的途中。",
	"Exception_Returned":                    "退件签收，寄件人已成功收到退件。",
	"Exception_NoBody":                      "找不到收件人，在派送之前发现的收件人信息异常。",
	"Exception_Security":                    "安全原因，在派送之前发现异常，包含安全、清关、费用问题。",
	"Exception_Damage":                      "损坏，在承运过程中发现货物损坏了。",
	"Exception_Rejected":                    "拒收，在派送之前接收到有收件人拒收情况。",
	"Exception_Delayed":                     "延误，因各种情况导致的可能超出原定的运输周期。",
	"Exception_Lost":                        "丢失，因各种情况导致的货物丢失。",
	"Exception_Destroyed":                   "销毁，因各种情况无法完成交付的货物并进行销毁。",
	"Exception_Cancel":                      "取消，因为各种情况物流订单被取消了。",
}

type _17trackEvent struct {
	Event _17trackEventType `json:"event,omitempty"` // 通知类型
	Data  struct {
		Number    string `json:"number,omitempty"`  // 物流单号
		Carrier   int    `json:"carrier,omitempty"` // 运输商代码
		Param     string `json:"param,omitempty"`   // 物流单号附加参数
		Tag       string `json:"tag,omitempty"`     // 自定义标签
		TrackInfo struct {
			ShippingInfo struct {
				ShipperAddress struct {
					Country     string `json:"country,omitempty"`     // 发件地国家或地区
					State       string `json:"state,omitempty"`       // 发件地州、省
					City        string `json:"city,omitempty"`        // 发件地城市
					Street      string `json:"street,omitempty"`      // 发件地街道
					PostalCode  string `json:"postal_code,omitempty"` // 发件地邮编
					Coordinates struct {
						Longitude string `json:"longitude,omitempty"` // 发件地经度
						Latitude  string `json:"latitude,omitempty"`  // 发件地纬度
					} `json:"coordinates,omitempty"` // 发件地位置坐标
				} `json:"shipper_address,omitempty"` // 发件地信息
				RecipientAddress struct {
					Country     string `json:"country,omitempty"`     // 收件地国家或地区
					State       string `json:"state,omitempty"`       // 收件地州、省
					City        string `json:"city,omitempty"`        // 收件地城市
					Street      string `json:"street,omitempty"`      // 收件地街道
					PostalCode  string `json:"postal_code,omitempty"` // 收件地邮编
					Coordinates struct {
						Longitude string `json:"longitude,omitempty"` // 收件地经度
						Latitude  string `json:"latitude,omitempty"`  // 收件地纬度
					} `json:"coordinates,omitempty"` // 收件地位置坐标
				} `json:"recipient_address,omitempty"` // 收件地信息
			} `json:"shipping_info,omitempty"` // 地区相关信息
			LatestStatus struct {
				Status         string `json:"status,omitempty"`           // 物流主状态
				SubStatus      string `json:"sub_status,omitempty"`       // 包裹子状态
				SubStatusDescr string `json:"sub_status_descr,omitempty"` // 状态描述
			} `json:"latest_status,omitempty"` // 最新状态
			LatestEvent struct {
				TimeIso string `json:"time_iso,omitempty"` // 事件发生时间（ISO 格式）
				TimeUtc string `json:"time_utc,omitempty"` // 事件发生时间（UTC 格式）
				TimeRaw struct {
					Date     string `json:"date,omitempty"`     // 年月日信息
					Time     string `json:"time,omitempty"`     // 时分秒信息
					Timezone string `json:"timezone,omitempty"` // 时区信息
				} `json:"time_raw,omitempty"` // 运输商提供的原始时间信息
				Description            string `json:"description,omitempty"` // 事件描述
				DescriptionTranslation struct {
					Lang        string `json:"lang,omitempty"`        // 翻译语言代码
					Description string `json:"description,omitempty"` // 翻译后的事件描述
				} `json:"description_translation,omitempty"` // 描述翻译
				Location  string `json:"location,omitempty"`   // 地点
				Stage     string `json:"stage,omitempty"`      // 里程碑状态
				SubStatus string `json:"sub_status,omitempty"` // 包裹子状态
				Address   struct {
					Country     string `json:"country,omitempty"`     // 国家地区
					State       string `json:"state,omitempty"`       // 州、省
					City        string `json:"city,omitempty"`        // 城市
					Street      string `json:"street,omitempty"`      // 街道
					PostalCode  string `json:"postal_code,omitempty"` // 邮编
					Coordinates struct {
						Longitude string `json:"longitude,omitempty"` // 经度
						Latitude  string `json:"latitude,omitempty"`  // 纬度
					} `json:"coordinates,omitempty"` // 位置坐标
				} `json:"address,omitempty"` // 地点信息
			} `json:"latest_event,omitempty"` // 最新事件
			TimeMetrics struct {
				DaysAfterOrder        int `json:"days_after_order,omitempty"`       // 运单时效
				DaysOfTransit         int `json:"days_of_transit,omitempty"`        // 运输时效
				DaysOfTransitDone     int `json:"days_of_transit_done,omitempty"`   // 妥投时效
				DaysAfterLastUpdate   int `json:"days_after_last_update,omitempty"` // 信息无更新天数
				EstimatedDeliveryDate struct {
					Source string `json:"source,omitempty"` // 时间区间提供者
					From   string `json:"from,omitempty"`   // 预计投递最早时间
					To     string `json:"to,omitempty"`     // 预计投递最晚时间
				} `json:"estimated_delivery_date,omitempty"` // 预期达到时间信息
			} `json:"time_metrics,omitempty"` // 时效相关信息
			Milestone []struct {
				KeyStage string `json:"key_stage,omitempty"` // 里程碑状态
				TimeIso  string `json:"time_iso,omitempty"`  // 事件发生时间（ISO 格式）
				TimeUtc  string `json:"time_utc,omitempty"`  // 事件发生时间（UTC 格式）
				TimeRaw  struct {
					Date     string `json:"date,omitempty"`     // 年月日信息
					Time     string `json:"time,omitempty"`     // 时分秒信息
					Timezone string `json:"timezone,omitempty"` // 时区信息
				} `json:"time_raw,omitempty"` // 运输商提供的原始时间信息
			} `json:"milestone,omitempty"` // 里程碑
			MiscInfo struct {
				RiskFactor      int    `json:"risk_factor,omitempty"`      // 包裹风险系数
				ServiceType     string `json:"service_type,omitempty"`     // 服务类型
				WeightRaw       string `json:"weight_raw,omitempty"`       // 原始重量信息
				WeightKg        string `json:"weight_kg,omitempty"`        // 转换为公斤的重量信息
				Pieces          string `json:"pieces,omitempty"`           // 件数
				Dimensions      string `json:"dimensions,omitempty"`       // 原始体积尺寸
				CustomerNumber  string `json:"customer_number,omitempty"`  // 收货客户单号
				ReferenceNumber string `json:"reference_number,omitempty"` // 参考号
				LocalNumber     string `json:"local_number,omitempty"`     // 尾程单号
				LocalProvider   string `json:"local_provider,omitempty"`   // 尾程运输商
				LocalKey        int    `json:"local_key,omitempty"`        // 尾程运输商代码
			} `json:"misc_info,omitempty"` // 包裹附属信息
			Tracking struct {
				ProvidersHash int `json:"providers_hash,omitempty"` // 哈希值
				Providers     []struct {
					Provider struct {
						Key      int    `json:"key,omitempty"`      // 运输商代码
						Name     string `json:"name,omitempty"`     // 运输商名称
						Alias    string `json:"alias,omitempty"`    // 运输商别名
						Tel      string `json:"tel,omitempty"`      // 运输商联系电话
						Homepage string `json:"homepage,omitempty"` // 运输商官网
						Country  string `json:"country,omitempty"`  // 运输商所属国家
					} `json:"provider,omitempty"` // 运输商信息
					ServiceType      string `json:"service_type,omitempty"`       // 服务类型
					LatestSyncStatus string `json:"latest_sync_status,omitempty"` // 最近同步状态
					LatestSyncTime   string `json:"latest_sync_time,omitempty"`   // 最近同步时间
					EventsHash       int    `json:"events_hash,omitempty"`        // 事件哈希值
					Events           []struct {
						TimeIso string `json:"time_iso,omitempty"` // 事件发生时间（ISO 格式）
						TimeUtc string `json:"time_utc,omitempty"` // 事件发生时间（UTC 格式）
						TimeRaw struct {
							Date     string `json:"date,omitempty"`     // 年月日信息
							Time     string `json:"time,omitempty"`     // 时分秒信息
							Timezone string `json:"timezone,omitempty"` // 时区信息
						} `json:"time_raw,omitempty"` // 运输商提供的原始时间信息
						Description            string `json:"description,omitempty"` // 事件描述
						DescriptionTranslation struct {
							Lang        string `json:"lang,omitempty"`        // 翻译语言代码
							Description string `json:"description,omitempty"` // 翻译后的事件描述
						} `json:"description_translation,omitempty"` // 描述翻译
						Location  string `json:"location,omitempty"`   // 地点
						Stage     string `json:"stage,omitempty"`      // 里程碑状态
						SubStatus string `json:"sub_status,omitempty"` // 包裹子状态
						Address   struct {
							Country     string `json:"country,omitempty"`     // 国家地区
							State       string `json:"state,omitempty"`       // 州、省
							City        string `json:"city,omitempty"`        // 城市
							Street      string `json:"street,omitempty"`      // 街道
							PostalCode  string `json:"postal_code,omitempty"` // 邮编
							Coordinates struct {
								Longitude string `json:"longitude,omitempty"` // 经度
								Latitude  string `json:"latitude,omitempty"`  // 纬度
							} `json:"coordinates,omitempty"` // 位置坐标
						} `json:"address,omitempty"` // 地点信息
					} `json:"events,omitempty"` // 事件集合
				} `json:"providers,omitempty"` // 运输商集合
			} `json:"tracking,omitempty"` // 物流信息
		} `json:"track_info,omitempty"` // 物流信息主结构节点
	} `json:"data,omitempty"` // 单号的跟踪详情
}
