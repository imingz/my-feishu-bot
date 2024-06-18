package qrcodeclient

import (
	"context"
	"encoding/json"
	"xiaoxiaojiqiren/internal/pkg/config"
)

func GetSatoken(ctx context.Context) (satoken string, err error) {
	var resp response
	err = client.Get("/web-app/auth/certificateLogin").
		SetQueryParam("openId", config.Get().Qrcode.OpenId).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", err
	}

	var data satokenData
	err = json.Unmarshal(resp.Data, &data)
	return data.Token, err
}

type satokenData struct {
	Account   string `json:"account"`
	Name      string `json:"name"`
	Token     string `json:"token"`
	TokenName string `json:"tokenName"`
	UserID    string `json:"userId"`
}
