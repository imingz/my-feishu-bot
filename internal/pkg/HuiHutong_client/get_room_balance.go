package huihutongclient

func (c *Client) getRoomId(buildingNumber, floorNumber, roomNumber string) (string, error) {
	var resp struct {
		Result []struct {
			RoomId     string `json:"roomId"`
			FangJianId string `json:"fangJianId"`
		} `json:"result"`
	}

	err := c.client.Get("/proxy/qy/sdcz/listRoom").
		SetQueryParams(map[string]string{
			"apartmentId": "2",
			"buildingId":  buildingNumber,
			"floorId":     floorNumber,
		}).
		Do().
		Into(&resp)

	if err != nil {
		return "", err
	}

	var trueRoomId string
	for _, data := range resp.Result {
		if data.RoomId == roomNumber {
			trueRoomId = data.FangJianId
			break
		}
	}

	return trueRoomId, nil
}

func (c *Client) GetRoomBalance(buildingNumber, floorNumber, roomNumber string) (float64, error) {
	roomId, err := c.getRoomId(buildingNumber, floorNumber, roomNumber) // TODO: 从多维表格获取
	if err != nil {
		return 0, err
	}

	var roomBalance struct {
		Result float64 `json:"result"`
	}
	err = c.client.Get("/proxy/qy/sdcz/getRoomBalance").
		SetQueryParams(map[string]string{
			"apartmentId": "7",
			"roomId":      roomId,
		}).
		Do().
		Into(&roomBalance)
	return roomBalance.Result, err
}
