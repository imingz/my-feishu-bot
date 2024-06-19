package hhtclient

import (
	"context"
	"encoding/json"
	"errors"
)

func GetQrcode(ctx context.Context) (qrcode string, err error) {
	satoken, err := GetSatoken(ctx)
	if err != nil {
		return "", err
	}

	var resp response
	err = client.Get("/pms/welcome/make-code-info").
		SetHeader("satoken", satoken).
		Do(ctx).
		Into(&resp)
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", errors.New(resp.Message)
	}

	var data makeCodeInfoData
	err = json.Unmarshal(resp.Data, &data)
	return data.QrCode, err
}

type makeCodeInfoData struct {
	Apartment    string `json:"apartment"`
	CompanyName  string `json:"companyName"`
	Locked       int64  `json:"locked"`
	LockedCount  int64  `json:"lockedCount"`
	LockedName   string `json:"lockedName"`
	MackCode     int64  `json:"mackCode"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	QrCode       string `json:"qrCode"`
	QrCodeStatus int64  `json:"qrCodeStatus"`
	Status       int64  `json:"status"`
	Text         string `json:"text"`
}
