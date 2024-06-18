package client

import (
	"context"
	"fmt"
	"io"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type ImageType string

const (
	ImageType_Message ImageType = "message" // 用于发送消息
	ImageType_Avatar  ImageType = "avatar"  // 用于设置头像
)

func (c *Client) Im_Image_Upload(imageType ImageType, image io.Reader) (string, error) {
	// 创建请求对象
	req := larkim.NewCreateImageReqBuilder().
		Body(larkim.NewCreateImageReqBodyBuilder().
			ImageType(string(imageType)).
			Image(image).
			Build()).
		Build()

	// 发起请求
	resp, err := c.client.Im.Image.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", err
	}

	// 服务端错误处理
	if !resp.Success() {
		return "", fmt.Errorf("client.Im.Image.Create failed, code: %d, msg: %s, log_id: %s", resp.Code, resp.Msg, resp.RequestId())
	}

	return *resp.Data.ImageKey, nil
}
