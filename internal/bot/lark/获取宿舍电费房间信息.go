package lark

import (
	"context"
	"errors"
	"fmt"
	"xiaoxiaojiqiren/internal/pkg/consts"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type 房间信息 struct {
	楼栋号 string
	楼层  string
	房间号 string
}

func (c *Client) 获取宿舍电费房间信息(ctx context.Context) (*房间信息, error) {
	event := ctx.Value(consts.KeyEvent).(*larkim.P2MessageReceiveV1Data)

	// 创建请求对象
	req := larkbitable.NewSearchAppTableRecordReqBuilder().
		AppToken(c.config.ConfigTable.AppToken).
		TableId(c.config.ConfigTable.TableId).
		UserIdType(`open_id`).
		PageSize(20).
		Body(larkbitable.NewSearchAppTableRecordReqBodyBuilder().
			FieldNames([]string{
				"宿舍楼栋号",
				"楼层",
				"房间号",
			}).
			Filter(larkbitable.NewFilterInfoBuilder().
				Conjunction(`and`).
				Conditions([]*larkbitable.Condition{
					larkbitable.NewConditionBuilder().
						FieldName(`提交人`).
						Operator(`is`).
						Value([]string{*event.Sender.SenderId.OpenId}).
						Build(),
				}).
				Build()).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Bitable.AppTableRecord.Search(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, err
	}

	// 服务端错误处理
	if !resp.Success() {
		return nil, errors.New(fmt.Sprintln(resp.Code, resp.Msg, resp.RequestId()))
	}

	// 业务处理
	return &房间信息{
		楼栋号: fmt.Sprintf("%.0f", resp.Data.Items[0].Fields["宿舍楼栋号"].(float64)),
		楼层:  fmt.Sprintf("%.0f", resp.Data.Items[0].Fields["楼层"].(float64)),
		房间号: fmt.Sprintf("%.0f", resp.Data.Items[0].Fields["房间号"].(float64)),
	}, err
}
