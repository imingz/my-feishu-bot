package lark

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func (c *Client) GetEnterpriseEmailByOpenId(openId string) (string, error) {
	// 创建请求对象
	req := larkcontact.NewGetUserReqBuilder().
		UserId(openId).
		Build()

	// 发起请求
	resp, err := c.client.Contact.User.Get(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", fmt.Errorf("请求失败, err: %v", err)
	}

	// 服务端错误处理
	if !resp.Success() {
		return "", fmt.Errorf("服务端错误, err: %v", larkcore.Prettify(resp))
	}

	// 业务处理
	if resp.Data.User.EnterpriseEmail == nil {
		return "", fmt.Errorf("用户未绑定企业邮箱, openId: %s", openId)
	}
	return *resp.Data.User.EnterpriseEmail, nil
}
