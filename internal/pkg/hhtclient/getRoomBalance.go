package hhtclient

import (
	"context"
	"errors"
)

func GetRoomBalance(ctx context.Context) (float64, error) {
	var resp getRoomBalanceResponse
	err := client.Get("/proxy/qy/sdcz/getRoomBalance?apartmentId=7&roomId=2290").
		Do(ctx).
		Into(&resp)

	if err != nil {
		return 0, err
	}

	if !resp.Success {
		return 0, errors.New(resp.Message)
	}

	return resp.Result, nil
}

type getRoomBalanceResponse struct {
	Code      int64       `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Result    float64     `json:"result"`
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
}
