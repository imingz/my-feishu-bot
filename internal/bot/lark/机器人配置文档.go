package lark

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// TODO: 暂时只支持最多20人
func (c *Client) Get宿舍电费配置() ([]*larkbitable.AppTableRecord, error) {
	// 创建请求对象
	req := larkbitable.NewSearchAppTableRecordReqBuilder().
		AppToken(c.config.ConfigTable.AppToken).
		TableId(c.config.ConfigTable.TableId).
		UserIdType(`open_id`).
		PageSize(20).
		Body(larkbitable.NewSearchAppTableRecordReqBodyBuilder().
			ViewId(c.config.ConfigTable.ViewId).
			FieldNames([]string{`宿舍楼栋号`, `房间号`, `提交人`, `楼层`, `电费阈值`}).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Bitable.AppTableRecord.Search(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, fmt.Errorf("请求失败, err: %v", err)
	}

	// 服务端错误处理
	if !resp.Success() {
		return nil, fmt.Errorf("服务端错误, err: %v", larkcore.Prettify(resp))
	}

	// 业务处理
	return resp.Data.Items, nil
}
