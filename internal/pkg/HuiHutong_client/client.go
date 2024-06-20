package huihutongclient

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req/v3"
)

type Client struct {
	client *req.Client
}

func NewClient() *Client {
	c := req.C().
		SetBaseURL("https://api.215123.cn").
		EnableDumpEachRequest().
		SetCommonErrorResult(&HuiHuTongResponse{}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				if dump := resp.Dump(); dump != "" {
					resp.Err = fmt.Errorf("%s\nraw content:\n%s", resp.Err.Error(), resp.Dump())
				}
				return nil
			}
			if err, ok := resp.ErrorResult().(*HuiHuTongResponse); ok {
				resp.Err = err
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("bad response, raw content:\n%s", resp.Dump())
				return nil
			}

			var data HuiHuTongResponse
			json.Unmarshal(resp.Bytes(), &data)
			if !data.Success {
				return &data
			}
			return nil
		})

	req.DevMode()

	return &Client{c}
}
